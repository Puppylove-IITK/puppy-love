package models

import (
	"context"

	"github.com/Puppylove-IITK/puppylove/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id      string `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Gender  string `json:"gender" bson:"gender"`
	Image   string `json:"image" bson:"image"`
	Pass    string `json:"passHash" bson:"passHash"`
	PrivK   string `json:"privKey" bson:"privKey"`
	PubK    string `json:"pubKey" bson:"pubKey"`
	AuthC   string `json:"authCode" bson:"authCode"`
	Data    string `json:"data" bson:"data"`
	Submit  bool   `json:"submitted" bson:"submitted"`
	Matches string `json:"matches" bson:"matches"`
	Vote    int    `json:"voted" bson:"voted"`
	Dirty   bool   `json:"dirty" bson:"dirty"`
	SPass   string `json:"savepass" bson:"savepass"`
}

// ----------------------------------------
type TypeUserNew struct {
	Id       string `json:"roll"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Image    string `json:"image"`
	PassHash string `json:"passHash"`
}

func NewUser(info *TypeUserNew) User {
	return User{
		Id:      info.Id,
		Name:    info.Name,
		Email:   info.Email,
		Gender:  info.Gender,
		Image:   info.Image,
		Pass:    info.PassHash,
		PrivK:   "",
		PubK:    "",
		AuthC:   utils.RandStringRunes(15),
		Data:    "",
		Submit:  false,
		Matches: "",
		Vote:    0,
		Dirty:   true,
		SPass:   "",
	}
}

// ----------------------------------------
type TypeUserFirst struct {
	Id       string `json:"roll"`
	AuthCode string `json:"authCode"`
	PassHash string `json:"passHash"`
	PubKey   string `json:"pubKey"`
	PrivKey  string `json:"privKey"`
	Data     string `json:"data"`
}

func (u User) FirstLogin(ctx context.Context, coll *mongo.Collection, info *TypeUserFirst) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": primitive.M{
			"authCode": "",
			"passHash": info.PassHash,
			"pubKey":   info.PubKey,
			"privKey":  info.PrivKey,
			"data":     info.Data,
			"matches":  "",
		},
	}

	return coll.UpdateOne(ctx, bson.M{"_id": info.Id}, update)
}

// ----------------------------------------
func (u User) ValidPass(pass string) bool {
	return u.Pass == pass
}

func (u User) SetField(field string, value interface{}) *options.FindOneAndUpdateOptions {
	return options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false).SetCollation(&options.Collation{})
}

type HeartsAndChoices struct {
	Hearts []GotHeart `json:"hearts"`
	Tokens Declare    `json:"tokens"`
}
