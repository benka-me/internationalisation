package rpc

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *App) GetMessage(ctx context.Context, req *internl.MessageRequest) (*internl.MessageResponse, error) {
	m := &internl.MessageResponse{}
	//filter := bson.D{{"data.username", }}
	_ = app.Client.Database(app.Collection.Database).
		Collection(req.GetCodeStr()).
		FindOne(context.TODO(), bson.D{}).Decode(&m)
	return &internl.MessageResponse{}, nil
}
