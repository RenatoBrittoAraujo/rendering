package main

import (
	"fmt"
	"os"

	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/driver"
	"github.com/renatobrittoaraujo/rendering/internal/logger"
)

func main() {
	configSource := &config.ConfigSource{
		SourceType: config.JSON,
		Address:    "config/tictactoe.json",
	}

	logger, err := logger.CreateLogger()
	if err != nil {
		fmt.Println("failed to initialize logger:", err)
		os.Exit(1)
	}

	config, err := config.LoadConfigFromSource(configSource, logger)
	if err != nil {
		fmt.Println("failed to initialize config:", err)
		os.Exit(1)
	}

	driver, err := driver.NewDriver(config, logger)
	if err != nil {
		fmt.Println("failed to initialize driver:", err)
		os.Exit(1)
	}

	err = driver.Run()
	if err != nil {
		fmt.Println("failed to run driver:", err)
		os.Exit(1)
	}
}
