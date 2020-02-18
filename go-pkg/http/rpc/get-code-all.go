package rpc

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"github.com/benka-me/internationalisation/go-pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Msg struct {
	Id primitive.ObjectID `bson:"omitempty"`
	Title string
	Code uint32
	Values map[uint32]string
}

//rpc GetCode(internl.CodeReq) returns (internl.Code);
//rpc GetAll(internl.AllReq) returns (internl.All);
func (app *App) GetCode(ctx context.Context, req *internl.CodeReq) (*internl.Code, error) {
	return mongo.GetCode(ctx, req, app.Client)
}
func (app *App) GetAll(ctx context.Context, req *internl.AllReq) (*internl.All, error) {
	return mongo.GetAll(ctx, req, app.Client)
}
