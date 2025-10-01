package commands

import (
	"fmt"
	"math/rand"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

func CommandCatch(cfg *config.AppConfig, client *api.Client, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing pokemon name")
	}

	url := ""
	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := client.GetPokemonStats(url, pokemonName)
	if err != nil {
		return fmt.Errorf("cannot catch: %v", err)
	}

	k := 500
	baseStat := pokemon.BaseExperience
	catchPossibility := float64(k) / float64(baseStat+k)

	if rand.Float64() < float64(catchPossibility) {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.Pokedex[pokemonName] = *pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
