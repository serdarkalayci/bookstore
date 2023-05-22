package mongodb

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

// HealthRepository represent a structure that will communicate to MongoDB to accomplish health related transactions
type HealthRepository struct {
	helper dbHelper
}

func newHealthRepository(client *mongo.Client, databaseName string) HealthRepository {
	return HealthRepository{
		helper: mongoHelper{coll: client.Database(databaseName).Collection(viper.GetString("BookInfoCollection"))},
	}
}

// Ready checks the arangodb connection
func (hr HealthRepository) Ready() bool {
	return false
}
