package main

import (
	"fmt"
	"os"

	"github.com/renatobrittoaraujo/rendering/internal/config"
	"github.com/renatobrittoaraujo/rendering/internal/driver"
	"github.com/renatobrittoaraujo/rendering/internal/logger"

	"github.com/fatih/color"
)

func main() {
	configSource := &config.ConfigSource{
		SourceType: config.JSON,
		Address:    "config/tictactoe.json",
	}

	logger, err := logger.CreateLogger()
	if err != nil {
		mainError("failed to initialize logger", err)
		os.Exit(1)
	}

	config, err := config.LoadConfigFromSource(configSource, logger)
	if err != nil {
		mainError("failed to initialize config", err)
		os.Exit(1)
	}

	driver, err := driver.NewDriver(config, logger)
	if err != nil {
		mainError("failed to initialize driver", err)
	}

	err = driver.Run()
	if err != nil {
		mainError("failed to run driver", err)
	}
}

func mainError(message string, err error) {
	fmt.Printf("\n")
	color.Red("[FATAL ERROR] %s; %s\n\n\n", message, fmt.Errorf("%w", err))
	os.Exit(1)
}
