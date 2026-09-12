package main

import (
	"context"
	"log"
	"ride-sharing/services/trip_service/internal/domain"
	"ride-sharing/services/trip_service/internal/infrastructure/repository"
	"ride-sharing/services/trip_service/internal/service"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInMemoryTripRepository()
	svc := service.NewService(inmemRepo)

	fare := domain.RideFareModel{
		ID:                primitive.NewObjectID(),
		UserID:            "user123",
		PackageSlug:       "van",
		TotalPriceInCents: 1500,
	}

	t, err := svc.CreateTrip(ctx, fare)

	if err != nil {
		log.Println("Error creating trip:", err)
	}

	log.Println("Trip created successfully:", t)

	for {
		time.Sleep(time.Second)
	}
}
