package cli

import (
	"os"
)

func commandExit(_ *Config, _ string) error {
	os.Exit(0)
	return nil
}
