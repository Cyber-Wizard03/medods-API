package database

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       uint   `bson:"_id" json:"id"`
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
}

func GetUser(filter2 string, filter1 interface{}) (User, error) {
	fil := bson.M{filter2: filter1}
	var user = User{}

	err := CollectionUser.FindOne(ctx, fil).Decode(&user)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	// This error means your query did not match any documents.
		// 	return
		// }
		panic(err)
	}

	return user, nil
}
func GetCheckUser() bool {
	fil := bson.M{}
	var user = User{}

	err := CollectionUser.FindOne(ctx, fil).Decode(&user)
	if err != nil {
		return false
	} else {
		return true
	}

}
