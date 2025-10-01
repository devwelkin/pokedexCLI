package commands

import (
	"fmt"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

func CommandInspect(cfg *config.AppConfig, client *api.Client, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing pokemon name")
	}

	pokemonName := args[0]

	if pokemon, exists := cfg.Pokedex[pokemonName]; exists {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Base Experience:", pokemon.BaseExperience)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}

		return nil
	}

	fmt.Printf("you didn't catch %s\n", pokemonName)
	return nil
}
