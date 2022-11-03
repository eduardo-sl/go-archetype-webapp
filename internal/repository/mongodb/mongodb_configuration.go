package mongodb

import (
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	connectionTimeout        = 60 * time.Second
	maxPoolSize       uint64 = 3000
)

type MongoDBConfiguration struct {
	DatabaseName string
	Uri          string
}

func NewMongoDBConfiguration() MongoDBConfiguration {
	return MongoDBConfiguration{
		DatabaseName: os.Getenv("MONGODB_NAME"),
		Uri:          os.Getenv("MONGODB_URI"),
	}
}

func (c MongoDBConfiguration) GetClientOptions() (options.ClientOptions, error) {

	clientOptions := options.ClientOptions{
		MaxPoolSize:    &maxPoolSize,
		ConnectTimeout: &connectionTimeout,
		ReadConcern:    readconcern.Local(),
	}
	clientOptions.ApplyURI(c.Uri)

	readPreferenceOptions := readpref.WithTags("use", "application")
	readPreference, err := readpref.New(readpref.SecondaryPreferredMode, readPreferenceOptions)

	if err != nil {
		return options.ClientOptions{}, err
	}

	clientOptions.SetReadPreference(readPreference)

	return clientOptions, nil
}
