package database

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type session struct {
	ID          primitive.ObjectID `bson:"_id"`
	ID_USER     uint               `bson:"id_user"`
	Fingerprint string             `bson:"fingerprint"`
	Date        int64              `bson:"date"`
	KeyOne      string             `bson:"keyone"`
	KeyTwo      string             `bson:"keytwo"`
}

func Generation() (a string) {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	length := 30
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
func Static() string {
	str := "yiPQBTyzDQVDeqcAknJb"
	return str
}
func CreateSession(id uint, keyone string, keytwo string, fingerprint string) {
	start := session{}
	currentTime := time.Now()
	TimeNow := currentTime.Add(3 * time.Hour).Unix()
	start = session{
		ID_USER:     id,
		Fingerprint: fingerprint,
		Date:        TimeNow,
		KeyOne:      keyone,
		KeyTwo:      keytwo,
	}
	Session.InsertOne(ctx, start)
}
func DropSession() {
	filter := bson.M{}
	Session.DeleteMany(ctx, filter)
}
func GetSession() session {
	data := session{}
	filter := bson.M{}
	err := Session.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
func UpdateSession(keyold string, keynew string) {
	item := GetSession()
	filter := bson.M{"_id": item.ID}
	fmt.Println(filter)
	update := bson.M{"$set": bson.M{
		"keyone": keynew,
	}}
	res, err := Session.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of items updated:", res.ModifiedCount)

}
