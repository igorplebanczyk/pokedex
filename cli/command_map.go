package cli

import (
	"fmt"
)

func commandMap(cfg *Config, _ string) error {
	resp, err := cfg.Client.GetLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Printf("-------------------------")
	fmt.Printf("\nLocation areas: \n")
	for _, locationArea := range resp.Results {
		println(locationArea.Name)
	}
	fmt.Printf("-------------------------\n")

	cfg.nextLocationAreasURL = resp.Next
	cfg.previousLocationAreasURL = resp.Previous

	return nil
}
