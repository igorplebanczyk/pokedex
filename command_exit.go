package main

import "os"

func commandExit(cfg *config, arg string) error {
	os.Exit(0)
	return nil
}
