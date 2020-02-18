package mongo

import (
	"context"
	"github.com/benka-me/internationalisation/go-pkg/internl"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestScene1(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	drop(client)

	tests := []struct {
		name    string

		newCode    *internl.NewCode
		newCodeWant    *internl.NewCodeRes
		newCodeErr bool

		setTitle *internl.SetTitleCode
		setTitleWant    *internl.SetTitleCodeRes
		setTitleErr bool

		setValue *internl.PushMessage
		setValueWant *internl.PushMessageRes
		setValueErr bool

		getMsg *internl.MessageRequest
		getMsgWant string
		getMsgErr bool
	}{
		{
			"Code 0",
			&internl.NewCode{
					Title: "Code 0",
					Code:  0,
			},
			&internl.NewCodeRes{},
			false,
			&internl.SetTitleCode{
				Title: "New Title Code 0",
				Code:  0,
			},
			&internl.SetTitleCodeRes{},
			false,
			&internl.PushMessage{
				Value: "Code 0 - EN value updated!!",
				Lang:  0,
				Code:  0,
			},
			&internl.PushMessageRes{},
			false,
			&internl.MessageRequest{
				ServiceRequester: "test",
				Lang:             0,
				Code:             0,
			},
			"Code 0 - EN value updated!!",
			false,
		},
		{
			"Code 0 dup",
			&internl.NewCode{
				Title: "Code 0 dup",
				Code:  0,
			},
			&internl.NewCodeRes{},
			true,
			&internl.SetTitleCode{
				Title: "Title Code 0 dup",
				Code:  0,
			},
			&internl.SetTitleCodeRes{},
			false,
			&internl.PushMessage{
				Value: "Code 0 dup - EN value updated!!",
				Lang:  0,
				Code:  0,
			},
			&internl.PushMessageRes{},
			false,
			&internl.MessageRequest{
				ServiceRequester: "test",
				Lang:             0,
				Code:             0,
			},
			"Code 0 dup - EN value updated!!",
			false,
		},
		{
			"Code 1",
			&internl.NewCode{
				Title: "Code 1",
				Code:  1,
			},
			&internl.NewCodeRes{},
			false,
			&internl.SetTitleCode{
				Title: "Title Code 1 updated",
				Code:  1,
			},
			&internl.SetTitleCodeRes{},
			false,
			&internl.PushMessage{
				Value: "Code 1 - FR valeur mise à jour",
				Lang:  1,
				Code:  1,
			},
			&internl.PushMessageRes{},
			false,
			&internl.MessageRequest{
				ServiceRequester: "test",
				Lang:             1,
				Code:             1,
			},
			"Code 1 - FR valeur mise à jour",
			false,
		},
	}
	const allLen = 2
	for _, tt := range tests {
		// Insert Code
		t.Run(tt.name + " insert code", func(t *testing.T) {
			if err := Insert(ctx, tt.newCode, client); (err != nil) != tt.newCodeErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.newCodeErr)
			}
		})

		// Set Title
		t.Run(tt.name + " set title", func(t *testing.T) {
			if err := SetCodeTitle(ctx, tt.setTitle, client); (err != nil) != tt.setTitleErr {
				t.Errorf("SetTitle() error = %v, wantErr %v", err, tt.setTitleErr)
			}
		})

		// Set Message
		//func UpdateValue(ctx context.Context, req *internl.PushMessage, client *mongo.Client) error {
		t.Run(tt.name + " set message", func(t *testing.T) {
			if err := UpdateValue(ctx, tt.setValue, client); (err != nil) != tt.setValueErr {
				t.Errorf("SetMessage() error = %v, wantErr %v", err, tt.setValueErr)
			}
		})

		// Get Message
		t.Run(tt.name + " get message", func(t *testing.T) {
			s, err := GetMessage(ctx, tt.getMsg, client)
			if (err != nil) != tt.getMsgErr{
				t.Errorf("GetMessage() error = %v, getMsgErr %v", err, tt.getMsgErr)
			} else if s != tt.getMsgWant {
				t.Errorf("GetMessage() error = %v, getMsgWant %v", s, tt.getMsgWant)
			}
		})
	}
		// GetAll
		t.Run("get all", func(t *testing.T) {
			all, err := GetAll(ctx, &internl.AllReq{}, client)
			if err != nil {
				t.Errorf("GetAll() error = %v", err)
			}
			if len(all.Codes) != allLen {
				t.Errorf("GetAll() error = bad len")
			}
		})
}
