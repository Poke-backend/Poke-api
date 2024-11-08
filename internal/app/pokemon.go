package app

import (
	"context"
	"fmt"
	"github.com/AmineZouitine/poke-backend/internal/app/aggregate"
)

type PokemonSvc struct {
	pkmRepo PokemonRepository
}

func NewPokemonSvc(pkmRepo PokemonRepository) *PokemonSvc {
	return &PokemonSvc{
		pkmRepo: pkmRepo,
	}
}

func (s *PokemonSvc) Create(ctx context.Context, pkm aggregate.Pokemon) (aggregate.Pokemon, error) {
	if err := pkm.Validate(); err != nil {
		return aggregate.Pokemon{}, fmt.Errorf("PokemonSvc.Create: %w", err)
	}

	return s.pkmRepo.Save(ctx, pkm)
}
