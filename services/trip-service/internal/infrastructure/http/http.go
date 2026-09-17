package http

import (
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
)

type HttpHandler struct {
	Service domain.TripService
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {

}

// fare := &domain.RideFareModel{
// 		ID:     primitive.NewObjectID(),
// 		UserID: "12",
// 	}

// t, err := svc.CreateTrip(ctx, fare)
// 	if err != nil {
// 		log.Println(err)
// 	}
