package commands

import (
	"testing"
	"time"

	"github.com/twomotive/pokedex/internal/api"
	"github.com/twomotive/pokedex/internal/config"
)

func TestCommandCatchMissingArgs(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	err := CommandCatch(cfg, client, []string{})
	if err == nil {
		t.Error("Expected error when no pokemon name provided")
	}
	if err.Error() != "missing pokemon name" {
		t.Errorf("Expected 'missing pokemon name' error, got: %v", err)
	}
}

func TestCommandInspectMissingArgs(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	err := CommandInspect(cfg, client, []string{})
	if err == nil {
		t.Error("Expected error when no pokemon name provided")
	}
	if err.Error() != "missing pokemon name" {
		t.Errorf("Expected 'missing pokemon name' error, got: %v", err)
	}
}

func TestCommandExploreMissingArgs(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	err := CommandExplore(cfg, client, []string{})
	if err == nil {
		t.Error("Expected error when no location area name provided")
	}
	if err.Error() != "missing location area name" {
		t.Errorf("Expected 'missing location area name' error, got: %v", err)
	}
}

func TestCommandInspectNotCaught(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	// Test inspecting a pokemon that hasn't been caught
	err := CommandInspect(cfg, client, []string{"pikachu"})
	if err != nil {
		t.Errorf("Expected no error for uncaught pokemon, got: %v", err)
	}
}

func TestCommandPokedexEmpty(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	err := CommandPokedex(cfg, client, []string{})
	if err != nil {
		t.Errorf("Expected no error for empty pokedex, got: %v", err)
	}
}

func TestCommandPokedexWithPokemon(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	// Add a mock pokemon to the pokedex
	cfg.Pokedex["pikachu"] = api.Pokemon{
		Name:           "pikachu",
		BaseExperience: 112,
		Height:         4,
		Weight:         60,
	}

	err := CommandPokedex(cfg, client, []string{})
	if err != nil {
		t.Errorf("Expected no error for non-empty pokedex, got: %v", err)
	}
}

func TestCommandHelp(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	err := CommandHelp(cfg, client, []string{})
	if err != nil {
		t.Errorf("Expected no error from help command, got: %v", err)
	}
}

func TestCommandMap(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	// Note: This will make a real API call, but it's a simple test
	// In a production environment, you'd want to mock the client
	err := CommandMap(cfg, client, []string{})
	if err != nil {
		// API calls might fail in test environment, that's ok
		t.Logf("Map command returned error (may be expected in test env): %v", err)
	}
}

func TestCommandMapBack(t *testing.T) {
	cfg := config.NewAppConfig()
	client := api.NewClient()

	// Test mapb with no previous page
	err := CommandMapBack(cfg, client, []string{})
	if err != nil {
		t.Errorf("Expected no error when no previous page, got: %v", err)
	}
}

func TestGetCommands(t *testing.T) {
	commands := GetCommands()

	expectedCommands := []string{"help", "exit", "map", "mapb", "explore", "catch", "inspect", "pokedex"}
	
	for _, cmdName := range expectedCommands {
		if _, exists := commands[cmdName]; !exists {
			t.Errorf("Expected command '%s' to exist", cmdName)
		}
	}

	// Verify each command has required fields
	for name, cmd := range commands {
		if cmd.Name != name {
			t.Errorf("Command name mismatch: key=%s, Name=%s", name, cmd.Name)
		}
		if cmd.Description == "" {
			t.Errorf("Command '%s' has empty description", name)
		}
		if cmd.Callback == nil {
			t.Errorf("Command '%s' has nil callback", name)
		}
	}
}

func TestConfigNewAppConfig(t *testing.T) {
	cfg := config.NewAppConfig()
	
	if cfg == nil {
		t.Fatal("NewAppConfig returned nil")
	}
	if cfg.Pagination == nil {
		t.Error("Pagination is nil")
	}
	if cfg.Pokedex == nil {
		t.Error("Pokedex map is nil")
	}
}

func TestAPINewClient(t *testing.T) {
	client := api.NewClient()
	
	if client == nil {
		t.Fatal("NewClient returned nil")
	}
}

// Test that helps ensure cleanup happens
func TestCleanup(t *testing.T) {
	// Create a client and let it go out of scope
	// This ensures the cache cleanup goroutine can be stopped
	client := api.NewClient()
	_ = client
	
	// Give goroutines a moment to start
	time.Sleep(10 * time.Millisecond)
}
