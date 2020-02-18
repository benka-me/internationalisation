package rpc

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"github.com/benka-me/internationalisation/go-pkg/mongo"
)


func (app *App) UpdateMessage(ctx context.Context, req *internl.PushMessage) (*internl.PushMessageRes, error) {
	return &internl.PushMessageRes{}, mongo.UpdateValue(ctx, req, app.Client)
}
