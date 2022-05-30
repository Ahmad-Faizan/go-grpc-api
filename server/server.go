package server

import (
	"context"
	"log"

	protos "github.com/Ahmad-faizan/go-grpc-api/protos/currency"
)

type Currency struct {
	log *log.Logger
}

func NewCurrency(l *log.Logger) *Currency {
	return &Currency{
		log: l,
	}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Println("Handle request for GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
