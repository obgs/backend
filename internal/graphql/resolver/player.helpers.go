package resolver

import (
	"context"

	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
)

func deleteRequestAndApprovals(ctx context.Context, tx *ent.Tx, requestID uuid.UUID) error {
	_, err := tx.PlayerSupervisionRequestApproval.Delete().Where(
		playersupervisionrequestapproval.HasSupervisionRequestWith(
			playersupervisionrequest.ID(requestID),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = tx.PlayerSupervisionRequest.Delete().Where(playersupervisionrequest.ID(requestID)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func addSupervisor(ctx context.Context, tx *ent.Tx, requestID uuid.UUID, supervisor *ent.User) error {
	err := deleteRequestAndApprovals(ctx, tx, requestID)
	if err != nil {
		return err
	}

	return tx.Player.Update().Where(
		player.HasSupervisionRequestsWith(
			playersupervisionrequest.ID(requestID),
		),
	).AddSupervisors(supervisor).Exec(ctx)
}

func handleSupervisionRequestApproval(ctx context.Context, tx *ent.Tx, requestID uuid.UUID, supervisor *ent.User) error {
	// approve the request
	_, err := tx.PlayerSupervisionRequestApproval.Update().Where(
		playersupervisionrequestapproval.HasSupervisionRequestWith(
			playersupervisionrequest.ID(requestID),
		),
	).SetApproved(true).Save(ctx)
	if err != nil {
		return err
	}

	// check if all approvals are done
	notApprovedCount, err := tx.PlayerSupervisionRequestApproval.Query().Where(
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
		return addSupervisor(ctx, tx, requestID, supervisor)
	}

	return nil
}
