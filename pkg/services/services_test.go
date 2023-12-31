package services

import (
	"goblin/config"
	"os"
	"testing"
)

var (
	c *Container
)

func TestMain(m *testing.M) {
	// Set test environment
	config.SwitchEnv(config.EnvTest)

	// Setup
	c = NewContainer()

	// TODO: DB creation and migration

	// Run tests
	exitCode := m.Run()

	// Teardown
	if err := c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}
