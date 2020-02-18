package rpc

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"github.com/benka-me/internationalisation/go-pkg/mongo"
)

func (app *App) InsertCode(ctx context.Context, req *internl.NewCode) (*internl.NewCodeRes, error) {
	return &internl.NewCodeRes{}, mongo.Insert(ctx, req, app.Client)
}
func (app *App) UpdateTitle(ctx context.Context, req *internl.SetTitleCode) (*internl.SetTitleCodeRes, error) {
	return &internl.SetTitleCodeRes{}, mongo.SetCodeTitle(ctx, req, app.Client)
}
