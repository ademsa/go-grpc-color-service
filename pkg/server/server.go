package server

import (
	"context"
	"github.com/ademsa/go-grpc-color-service/pkg/service"
	"log"
)

// ColorServer for Color Service
type ColorServer struct {
	service.UnimplementedColorServiceServer
}

// GetColor implementation
func (s *ColorServer) GetColor(c context.Context, r *service.GetColorRequest) (*service.GetColorResponse, error) {
	log.Print("Incoming GetColor request")
	return &service.GetColorResponse{Color: "2979FF"}, nil
}
