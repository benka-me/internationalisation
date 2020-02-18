package mongo

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

const (
	db = "internationalisation"
	messages = "messages"
)

type Msg struct {
	Id primitive.ObjectID `bson:"omitempty"`
	Title string
	Code uint32
	Values map[uint32]string
}

func Insert(ctx context.Context, req *internl.NewCode, client *mongo.Client) error {
	filter := bson.D{{"code", req.GetCode()}}
	err := client.Database(db).Collection(messages).FindOne(ctx, filter).Err()
	if err == nil {
		return status.Error(codes.AlreadyExists, "")
	}

	_, _ = client.Database(db).Collection(messages).
		InsertOne(ctx, req)

	return nil
}

func GetMessage(ctx context.Context, req *internl.MessageRequest, client *mongo.Client) (string, error) {
	c := &internl.Code{}
	filter := bson.D{{"code", req.GetCode()}}

	err := client.Database(db).Collection(messages).
		FindOne(ctx, filter).Decode(c)
	if err != nil {
		return "", err
	}
	return c.Values[uint32(req.Lang)], nil
}

func UpdateValue(ctx context.Context, req *internl.PushMessage, client *mongo.Client) error {
	lg := strconv.FormatInt(int64(req.Lang), 10)
	filter := bson.D{{"code", req.Code}}
	update := bson.D{{"$set",
		bson.D{{"values." + lg, req.Value}},
	}}

	_, err := client.Database(db).Collection(messages).
		UpdateOne(ctx, filter, update)

	return err
}

func SetCodeTitle(ctx context.Context, req *internl.SetTitleCode, client *mongo.Client) error {
	filter := bson.D{{"code", req.Code}}
	update := bson.D{{"$set",
		bson.D{{"title", req.Title}},
	}}

	_, err := client.Database(db).Collection(messages).
		UpdateOne(ctx, filter, update)

	return err
}

func GetCode(ctx context.Context, req *internl.CodeReq, client *mongo.Client) (*internl.Code, error) {
	m := &internl.Code{}
	filter := bson.D{{"code", req.GetCode()}}

	_ = client.Database(db).Collection(messages).
		FindOne(ctx, filter).Decode(m)

	return m, nil
}

func GetAll(ctx context.Context, req *internl.AllReq, client *mongo.Client) (*internl.All, error) {
	all := &internl.All{}
	filter := bson.D{}
	all.Values = make(map[uint32]string)

	cur, _ := client.Database(db).Collection(messages).
		Find(ctx, filter)

	for cur.Next(ctx) {
		tmp := struct {
			Title string
			Code uint32
		}{}
		err := cur.Decode(&tmp)
		if err != nil {
			return all, err
		}
		all.Values[tmp.Code] = tmp.Title
	}

	return all, nil
}
