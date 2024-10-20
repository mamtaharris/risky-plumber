package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type PaginationConfig struct {
	Limit  int `json:"PAGINATION_LIMIT"`
	Offset int `json:"PAGINATION_OFFSET"`
}

var Pagination *PaginationConfig

func loadPaginationDefaultConfig() {
	Pagination = &PaginationConfig{}
	err := envconfig.Process("pagination", Pagination)
	if err != nil {
		log.Fatal(err.Error())
	}
}
