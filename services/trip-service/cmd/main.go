package main

import (
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)

	_ = svc
}
