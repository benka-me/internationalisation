package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func drop(client *mongo.Client) {
	client.Database(db).Collection(messages).Drop(context.TODO())
}
