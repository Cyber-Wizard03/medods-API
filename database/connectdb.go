package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type New struct {
	ID      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Image   string             `bson:"image"`
	Text    string             `bson:"text"`
	Date    string             `bson:"date"`
	Content []Item             `bson:"content"`
}
type Item struct {
	Type    string   `bson:"type"`
	Content []string `bson:"content"`
}
type Newj struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Image   string  `json:"image"`
	Text    string  `json:"text"`
	Date    string  `json:"date"`
	Content []Itemj `json:"content"`
}
type Itemj struct {
	Type    string   `json:"type"`
	Content []string `json:"content"`
}
type Contactsj struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
type Contacts struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Number string             `bson:"number"`
}

type Documents struct {
	ID       primitive.ObjectID `bson:"_id"`
	Search   string             `bson:"search" json:"search"`
	Resource string             `bson:"resource" json:"resource"`
	Name     string             `bson:"name" json:"name"`
}
type Useful struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Url   string             `bson:"url" json:"url"`
	Name  string             `bson:"name" json:"name"`
	Image string             `bson:"image" json:"image"`
}
type Services struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Type  string             `bson:"type" json:"type"`
	Image []string           `bson:"image" json:"image"`
}
type Personal struct {
	ID            primitive.ObjectID `bson:"_id"`
	Img           string             `bson:"img"`
	Surname       string             `bson:"surname"`
	Name          string             `bson:"name"`
	Patronymic    string             `bson:"patronymic"`
	Department    string             `bson:"department"`
	IdDepartment  string             `bson:"iddepartment"`
	Cabinet       string             `bson:"cabinet"`
	Post          string             `bson:"post"`
	Education     string             `bson:"education"`
	Certificate   string             `bson:"certificate"`
	Qualification string             `bson:"qualification"`
}
type Branch struct {
	ID       primitive.ObjectID `bson:"_id"`
	IdBranch string             `json:"idbranch"`
	Branch   string             `bson:"branch"`
	Name     string             `bson:"name"`
}
type PersonalJ struct {
	ID            string `json:"id"`
	Img           string `json:"img"`
	Surname       string `json:"surname"`
	Name          string `json:"name"`
	Patronymic    string `json:"patronymic"`
	Department    string `json:"department"`
	IdDepartment  string `json:"iddepartment"`
	Cabinet       string `json:"cabinet"`
	Post          string `json:"post"`
	Education     string `json:"education"`
	Certificate   string `json:"certificate"`
	Qualification string `json:"qualification"`
}
type BranchJ struct {
	ID       string `json:"id"`
	IdBranch string `json:"idbranch"`
	Branch   string `json:"branch"`
	Name     string `json:"name"`
}
type Vacancy struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Profession string             `bson:"profession" json:"profession"`
	Experience string             `bson:"experience" json:"experience"`
	Education  string             `bson:"education" json:"education"`
	Salary     string             `bson:"salary" json:"salary"`
	Schedule   string             `bson:"schedule" json:"schedule"`
	Other      string             `bson:"other" json:"other"`
}
type Info struct {
	ID     primitive.ObjectID `bson:"_id"`
	Search string             `bson:"search" json:"search"`
	Text   string             `bson:"text" json:"text"`
}
type MesFeedback struct {
	ID           primitive.ObjectID `bson:"_id"`
	FIO          string             `bson:"FIO" json:"FIO"`
	Phone        string             `bson:"Phone" json:"Phone"`
	Organization string             `bson:"Organization" json:"Organization"`
	Department   string             `bson:"Department" json:"Department"`
	Doctor       string             `bson:"Doctor" json:"Doctor"`
	Type_appeal  string             `bson:"Type_appeal" json:"Type_appeal"`
	Messages     string             `bson:"Messages" json:"Messages"`
}

var (
	CollectionCon        *mongo.Collection
	CollectionNew        *mongo.Collection
	CollectionPrevention *mongo.Collection
	CollectionContacts   *mongo.Collection
	CollectionDocuments  *mongo.Collection
	CollectionUser       *mongo.Collection
	Session              *mongo.Collection
	CollectionServices   *mongo.Collection
	CollectionPersonal   *mongo.Collection
	CollectionBranch     *mongo.Collection
	CollectionVacancy    *mongo.Collection
	CollectionInfo       *mongo.Collection
	CollectionFeedback   *mongo.Collection
)
var ctx = context.TODO()

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	CollectionCon = client.Database("guzkod").Collection("Сontact")
	CollectionNew = client.Database("guzkod").Collection("News")
	CollectionPrevention = client.Database("guzkod").Collection("Prevention")
	CollectionContacts = client.Database("guzkod").Collection("Contacts")
	CollectionDocuments = client.Database("guzkod").Collection("Documents")
	CollectionUser = client.Database("guzkod").Collection("User")
	Session = client.Database("guzkod").Collection("Session")
	CollectionServices = client.Database("guzkod").Collection("Services")
	CollectionPersonal = client.Database("guzkod").Collection("Personal")
	CollectionBranch = client.Database("guzkod").Collection("Branch")
	CollectionVacancy = client.Database("guzkod").Collection("Vacancy")
	CollectionInfo = client.Database("guzkod").Collection("Info")
	CollectionFeedback = client.Database("Feedback").Collection("Messages")
}
