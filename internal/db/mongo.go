package db

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Log struct {
	LogMessage string `bson:"log_message"`
	Time       time.Time
	IpAdress   string
	RemoteAddr string
}

func ConnectMongoSaveLog(logData string, c echo.Context) error {
	// MongoDB bağlantısı oluşturma
	clientOptions := options.Client().ApplyURI("mongodb://45.12.81.218:27021/?directConnection=true")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// MongoDB'ye ping atarak bağlantıyı test etme
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// "logs" collection'ını oluşturma veya varsa alma
	db := client.Database("logdb")
	collection := db.Collection("logs")

	// Collection'ın var olup olmadığını kontrol etme
	if _, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"log_id": 1},
			Options: options.Index().SetUnique(false),
		},
	); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Collection 'logs' already exists")
		} else {
			return err
		}
	} else {
		log.Println("Collection 'logs' created")
	}

	logEntry := Log{LogMessage: logData, Time: time.Now(), IpAdress: c.RealIP(), RemoteAddr: c.Request().RemoteAddr}
	res, err := collection.InsertOne(context.Background(), logEntry)
	if err != nil {
		return err
	}
	log.Printf("Inserted %v documents into 'logs' collection\n", res.InsertedID)
	return nil
}
