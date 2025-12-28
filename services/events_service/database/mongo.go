package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	querries *mongo.Client
}

func NewMongoDb(client *mongo.Client)*MongoDB{
	return &MongoDB{
		querries:client,
	}
}

func ConnectMongo()(*mongo.Client,context.CancelFunc) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping failed:", err)
	}

	
	log.Println("✅ MongoDB connected")

	return client,cancel
}

// GenerateEvent(ctx context.Context,event *models.Event)(*models.Event,error)
// DeleteEvent(ctx context.Context,eventId any)(error)
// UpdateEvent(ctx context.Context,event *models.Event)(models.Event,error)
// GetEvent(ctx context.Context,eventId any)(models.Event,error)

func (db *MongoDB)GenerateEvent(ctx context.Context,event *models.Event)(*models.Event,error){
	return nil,nil
}

func (db *MongoDB)DeleteEvent(ctx context.Context,eventId any)(error){
	return nil
}

func (db *MongoDB)UpdateEvent(ctx context.Context,event *models.Event)(models.Event,error){
	return models.Event{},nil
}

func (db *MongoDB)GetEvent(ctx context.Context,eventId any)(models.Event,error){
	return models.Event{},nil
}