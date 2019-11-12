package main

import "github.io/whitnat/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	var ps = spaceadventure.PlanetarySystem{Name: "Solare System"}
	spaceadventure.Start(ps)
}
