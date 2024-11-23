package main

import (
	"neuro-most/category-service/config"
	"neuro-most/category-service/internal/infra"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	infra.Config(cfg).Database().Serve().Start()
}
