package resolver

import (
	"context"

	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

func deleteRequestAndApprovals(ctx context.Context, client *ent.Client, requestID uuid.UUID) error {
	_, err := client.PlayerSupervisionRequestApproval.Delete().Where(
		playersupervisionrequestapproval.HasSupervisionRequestWith(
			playersupervisionrequest.ID(requestID),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = client.PlayerSupervisionRequest.Delete().Where(playersupervisionrequest.ID(requestID)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func addSupervisor(ctx context.Context, client *ent.Client, requestID uuid.UUID) error {
	request, err := client.PlayerSupervisionRequest.Get(ctx, requestID)
	if err != nil {
		return err
	}

	supervisor, err := request.Sender(ctx)
	if err != nil {
		return err
	}

	err = client.Player.Update().Where(
		player.HasSupervisionRequestsWith(
			playersupervisionrequest.ID(requestID),
		),
	).AddSupervisors(supervisor).Exec(ctx)
	if err != nil {
		return err
	}

	return deleteRequestAndApprovals(ctx, client, requestID)
}

func handleSupervisionRequestApproval(ctx context.Context, client *ent.Client, approverID, requestID uuid.UUID) error {
	// approve the request
	_, err := client.PlayerSupervisionRequestApproval.Update().Where(
		playersupervisionrequestapproval.HasSupervisionRequestWith(
			playersupervisionrequest.ID(requestID),
		),
		playersupervisionrequestapproval.HasApproverWith(
			user.ID(approverID),
		),
	).SetApproved(true).Save(ctx)
	if err != nil {
		return err
	}

	// check if all approvals are done
	notApprovedCount, err := client.PlayerSupervisionRequestApproval.Query().Where(
		playersupervisionrequestapproval.HasSupervisionRequestWith(
			playersupervisionrequest.ID(requestID),
		),
		playersupervisionrequestapproval.Or(
			playersupervisionrequestapproval.Approved(false),
			playersupervisionrequestapproval.ApprovedIsNil(),
		),
	).Count(ctx)
	if err != nil {
		return err
	}

	// if all approvals are done, we add the new supervisor and delete the request
	if notApprovedCount == 0 {
		return addSupervisor(ctx, client, requestID)
	}

	return nil
}
