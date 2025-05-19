package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	rootCtx := context.Background()
	mongoDbUri := os.Getenv("MONGO_DB_URI")
	if mongoDbUri == "" {
		logger.Fatal().Msg("MONGO_DB_URI is not set")
	}

	db, err := mongo.Connect(rootCtx, options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to mongo db")
	}

	err = db.Ping(rootCtx, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to ping mongo db")
	}

	logger.Info().Msg("connected to mongo db")
}
