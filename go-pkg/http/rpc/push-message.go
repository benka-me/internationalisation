package rpc

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
)

func (app *App) PushMessage(ctx context.Context, req *internl.PushMessageReq) (*internl.PushMessageRes, error) {
	_, _ = app.Client.Database(app.Collection.Database).
		Collection(req).
		InsertOne(context.TODO(), req.PushMessage)
	return &internl.PushMessageRes{}, nil
}
