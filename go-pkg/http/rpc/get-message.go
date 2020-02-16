package rpc

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *App) GetMessage(ctx context.Context, req *internl.MessageRequest) (*internl.MessageResponse, error) {
	m := &internl.Message{}
	filter := bson.D{{"data.username", identifier.GetName()},}
	app.Client.Database(app.Collection.Database).Collection(app.Collection.Messages).
		FindOne(context.TODO(), filter).Decode(&u)
	return &internl.MessageResponse{}, nil
}
