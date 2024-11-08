package app

import (
	"context"
	"github.com/AmineZouitine/poke-backend/internal/app/aggregate"
)

type PokemonRepository interface {
	Save(ctx context.Context, pkm aggregate.Pokemon) (aggregate.Pokemon, error)
	ListPokemon(ctx context.Context, pkm aggregate.Pokemon) ([]aggregate.Pokemon, error)
	PokemonByID(ctx context.Context, id string) (aggregate.Pokemon, error)
}
