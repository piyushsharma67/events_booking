package que

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/piyushsharma67/events_booking/services/booking_service/domain"
	"github.com/piyushsharma67/events_booking/services/booking_service/service"
	"github.com/streadway/amqp"
)

func GenerateSeatsHandler(
	bookingService *service.SeatsService,
) func(amqp.Delivery) error {

	return func(d amqp.Delivery) error {
		var msg GenerateSeatsMessage

		// 1️⃣ Parse message
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			log.Println("invalid generate seats message:", err)
			return err // NACK → retry or DLQ
		}

		if msg.EventID == "" {
			return errors.New("event_id is empty")
		}

		// 2️⃣ Build domain seats
		seats := make([]domain.Seat, 0)

		for _, r := range msg.SeatLayout.Rows {
			for i := 1; i <= r.Count; i++ {
				seats = append(seats, domain.Seat{
					EventID:    msg.EventID,
					RowID:      r.Row,
					SeatNumber: i,
					Status:     "AVAILABLE",
				})
			}
		}

		if len(seats) == 0 {
			log.Println("no seats to generate for event:", msg.EventID)
			return nil // ACK safely
		}

		// 3️⃣ Call service layer
		ctx := context.Background()
		if err := bookingService.GenerateSeats(ctx, seats); err != nil {
			log.Println("generate seats failed:", err)
			return err
		}

		log.Printf(
			"generated %d seats for event %s",
			len(seats),
			msg.EventID,
		)

		return nil // ACK
	}
}
