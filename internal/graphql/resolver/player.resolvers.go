package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
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

	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	player, err := tx.Player.Get(ctx, input.PlayerID)
	if err != nil {
		return nil, err
	}

	currentSupervisors, err := tx.Player.QuerySupervisors(player).All(ctx)
	if err != nil {
		return nil, err
	}

	request, err := tx.PlayerSupervisionRequest.Create().SetPlayer(player).SetSender(user).SetMessage(*input.Message).Save(ctx)
	if err != nil {
		return nil, err
	}

	approvals := make([]*ent.PlayerSupervisionRequestApprovalCreate, len(currentSupervisors))
	for i, supervisor := range currentSupervisors {
		approvals[i] = tx.PlayerSupervisionRequestApproval.Create().
			SetApprover(supervisor).
			SetSupervisionRequest(request)
	}

	_, err = tx.PlayerSupervisionRequestApproval.CreateBulk(approvals...).Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
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

	tx, err := r.client.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}

	// if the request is rejected, we remove all approvals and delete the request
	if !input.Approved {
		err = deleteRequestAndApprovals(ctx, tx, input.RequestID)
	} else {
		err = handleSupervisionRequestApproval(ctx, tx, input.RequestID, u)
	}
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return false, err
	}

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
