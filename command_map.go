package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	res, err := cfg.client.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	cfg.next = res.Next
	cfg.prev = res.Prev
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prev == "" {
		return errors.New("no previous url to go")
	}
	res, err := cfg.client.GetLocationAreas(cfg.prev)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	cfg.next = res.Next
	cfg.prev = res.Prev
	return nil
}
