package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// CreatePlayer is the resolver for the createPlayer field.
func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.CreatePlayerInput) (*ent.Player, error) {
	supervisor, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return ent.FromContext(ctx).Player.Create().SetName(input.Name).AddSupervisors(supervisor).Save(ctx)
}

// RequestPlayerSupervision is the resolver for the requestPlayerSupervision field.
func (r *mutationResolver) RequestPlayerSupervision(ctx context.Context, input *model.RequestPlayerSupervisionInput) (*ent.PlayerSupervisionRequest, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	p, err := r.client.Player.Query().Where(
		player.ID(input.PlayerID),
	).WithOwner().Only(ctx)
	if err != nil {
		return nil, err
	}

	currentSupervisors, err := r.client.Player.QuerySupervisors(p).All(ctx)
	if err != nil {
		return nil, err
	}

	request, err := r.client.PlayerSupervisionRequest.Create().SetPlayer(p).SetSender(user).SetMessage(*input.Message).Save(ctx)
	if err != nil {
		return nil, err
	}

	owner := p.Edges.Owner
	approvalCount := len(currentSupervisors)
	if owner != nil {
		approvalCount++
	}

	approvals := make([]*ent.PlayerSupervisionRequestApprovalCreate, approvalCount)
	for i, supervisor := range currentSupervisors {
		approvals[i] = r.client.PlayerSupervisionRequestApproval.Create().
			SetApprover(supervisor).
			SetSupervisionRequest(request)
	}

	if owner != nil {
		approvals[approvalCount-1] = r.client.PlayerSupervisionRequestApproval.Create().
			SetApprover(owner).
			SetSupervisionRequest(request)
	}

	_, err = r.client.PlayerSupervisionRequestApproval.CreateBulk(approvals...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return request, nil
}

// ResolvePlayerSupervisionRequest is the resolver for the resolvePlayerSupervisionRequest field.
func (r *mutationResolver) ResolvePlayerSupervisionRequest(ctx context.Context, input model.ResolvePlayerSupervisionRequestInput) (bool, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return false, err
	}

	// if the request is rejected, we remove all approvals and delete the request
	if !input.Approved {
		err = deleteRequestAndApprovals(ctx, r.client, input.RequestID)
	} else {
		err = handleSupervisionRequestApproval(ctx, r.client, u.ID, input.RequestID)
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
