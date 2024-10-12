package cli

import (
	"fmt"
)

func commandMapb(cfg *Config, _ string) error {
	if cfg.previousLocationAreasURL == nil {
		println("You're already at the beginning of the list.")
		return nil
	}

	locationAreasResp, err := cfg.Client.GetLocationAreas(cfg.previousLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Printf("-------------------------")
	fmt.Printf("\nLocation areas: \n")
	for _, locationArea := range locationAreasResp.Results {
		println(locationArea.Name)
	}
	fmt.Printf("-------------------------\n")

	cfg.nextLocationAreasURL = locationAreasResp.Next
	cfg.previousLocationAreasURL = locationAreasResp.Previous

	return nil
}
