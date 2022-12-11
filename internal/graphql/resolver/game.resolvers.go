package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/gamefavorite"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// Favorites is the resolver for the favorites field.
func (r *gameResolver) Favorites(ctx context.Context, obj *ent.Game) (*model.Favorites, error) {
	total, err := obj.QueryFavorites().Count(ctx)
	if err != nil {
		return nil, err
	}

	favs, err := obj.QueryFavorites().WithUser().Limit(MAX_FAVS).All(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*ent.User, 0, len(favs))
	for _, fav := range favs {
		users = append(users, fav.Edges.User)
	}

	return &model.Favorites{
		Total: total,
		Users: users,
	}, nil
}

// IsFavorite is the resolver for the isFavorite field.
func (r *gameResolver) IsFavorite(ctx context.Context, obj *ent.Game) (bool, error) {
	u, _ := auth.UserFromContext(ctx)
	if u == nil {
		return false, nil
	}

	return obj.QueryFavorites().Where(gamefavorite.HasUserWith(user.ID(u.ID))).Exist(ctx)
}

// CreateGame is the resolver for the createGame field.
func (r *mutationResolver) CreateGame(ctx context.Context, input model.CreateGameInput) (*ent.Game, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	createDescriptons := make([]*ent.StatDescriptionCreate, 0, len(input.StatDescriptions))
	for _, desc := range input.StatDescriptions {
		create := r.client.StatDescription.Create().
			SetType(desc.Type).
			SetName(desc.Name).
			SetDescription(*desc.Description)
		switch desc.Type {
		case stat.Numeric:
			if desc.EnumStatInput != nil {
				return nil, fmt.Errorf("enumStatInput is not allowed for numeric stat")
			}
		case stat.Enum:
			if desc.EnumStatInput == nil || desc.EnumStatInput.PossibleValues == nil {
				return nil, fmt.Errorf("possible values are required for enum stat")
			}
			if len(desc.EnumStatInput.PossibleValues) < MIN_ENUM_VALUES {
				return nil, fmt.Errorf("at least %d possible values are required for enum stat", MIN_ENUM_VALUES)
			}
			create.SetPossibleValues(desc.EnumStatInput.PossibleValues)
		}
		createDescriptons = append(
			createDescriptons,
			create,
		)
	}

	descriptions, err := r.client.StatDescription.CreateBulk(createDescriptons...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.client.Game.Create().
		SetName(input.Name).
		SetDescription(*input.Description).
		SetBoardgamegeekURL(*input.BoardgamegeekURL).
		SetAuthor(u).
		SetMinPlayers(input.MinPlayers).
		SetMaxPlayers(input.MaxPlayers).
		AddStatDescriptions(descriptions...).
		Save(ctx)
}

// AddOrRemoveGameFromFavorites is the resolver for the addOrRemoveGameFromFavorites field.
func (r *mutationResolver) AddOrRemoveGameFromFavorites(ctx context.Context, gameID guidgql.GUID, favorite bool) (bool, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return false, err
	}

	if favorite {
		_, err = r.client.GameFavorite.Create().
			SetUser(u).
			SetGameID(gameID).
			Save(ctx)
	} else {
		_, err = r.client.GameFavorite.Delete().Where(
			gamefavorite.And(
				gamefavorite.HasUserWith(user.ID(u.ID)),
				gamefavorite.HasGameWith(game.ID(gameID)),
			),
		).Exec(ctx)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
