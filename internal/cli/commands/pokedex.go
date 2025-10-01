package commands

import (
	"fmt"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

func CommandPokedex(cfg *config.AppConfig, client *api.Client, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.Pokedex {
		fmt.Printf(" - %s \n", name)
	}

	return nil
}
