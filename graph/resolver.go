package graph

import "github.com/Am1rArsalan/kelvin-green/service"

type Resolver struct {
	service service.ServiceI
}

func NewResolver(service service.ServiceI) *Resolver {
	return &Resolver{
		service: service,
	}
}
