package mongodb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMongoDBConfiguration(t *testing.T) {
	t.Setenv("MONGODB_NAME", "database-name")
	t.Setenv("MONGODB_URI", "mongodb://localhost:27027")
	mongoDBConfiguration := NewMongoDBConfiguration()

	assert.Equal(t, mongoDBConfiguration.DatabaseName, "database-name", "The database names should be the same.")
	assert.Equal(t, mongoDBConfiguration.Uri, "mongodb://localhost:27027", "The uris should be the same.")
}

func TestMongoDBConfigurationOptions(t *testing.T) {
	expectedMaxPoolSize := uint64(3000)
	expectedConnectTimeout := 60 * time.Second
	options, err := NewMongoDBConfiguration().GetClientOptions()

	assert.Nil(t, err)

	maxPoolSize := *options.MaxPoolSize
	connectTimeout := *options.ConnectTimeout
	assert.Equal(t, expectedMaxPoolSize, maxPoolSize)

	assert.Equal(t, expectedConnectTimeout, connectTimeout)
}
