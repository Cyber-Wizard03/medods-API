package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"

	"guzkod/database"
)

func NewInfo(c *gin.Context) {
	var input = database.Info{}
	input.ID = primitive.NewObjectID()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionInfo)
}
func FeedbackPost(c *gin.Context) {
	var input = database.MesFeedback{}
	input.ID = primitive.NewObjectID()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionFeedback)
}
func FeedbackGet(c *gin.Context) {
	a := bson.M{}
	var res, _ = database.GetFeedback(a, database.CollectionFeedback)
	// res = Reverse(res)
	var f = database.MesFeedback{}
	var pull = []database.MesFeedback{}
	for _, b := range res {
		// var stringObjectID = b.ID.Hex()

		f = database.MesFeedback{
			// ID:           stringObjectID,
			FIO:          b.FIO,
			Phone:        b.Phone,
			Organization: b.Organization,
			Department:   b.Department,
			Doctor:       b.Doctor,
			Type_appeal:  b.Type_appeal,
			Messages:     b.Messages,
		}
		pull = append(pull, f)

	}
	c.IndentedJSON(http.StatusOK, pull)
}
func getInfo(c *gin.Context) {
	key := c.Param("Doc")
	a := bson.M{"search": key}
	pull := []database.Info{}
	var res, _ = database.GetInfo(a, database.CollectionInfo)
	pull = append(pull, res)
	c.IndentedJSON(http.StatusOK, pull)
}

func NewVacancy(c *gin.Context) {
	var input = database.Vacancy{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = primitive.NewObjectID()
	database.CreateTask(input, database.CollectionVacancy)
}

func VacancyF(c *gin.Context) {
	a := bson.M{}
	var res, _ = database.GetVacancy(a, database.CollectionVacancy)

	c.IndentedJSON(http.StatusOK, res)
}
func Person(c *gin.Context) {
	key := c.Param("IDperson")
	objID, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		panic(err)
	}
	a := bson.M{"_id": objID}
	var f = database.PersonalJ{}
	var pull = []database.PersonalJ{}
	var res, _ = database.GetPer(a, database.CollectionPersonal)
	for _, b := range res {
		var stringObjectID = b.ID.Hex()
		f = database.PersonalJ{
			ID:            stringObjectID,
			Img:           b.Img,
			Surname:       b.Surname,
			Name:          b.Name,
			Patronymic:    b.Patronymic,
			Department:    b.Department,
			Cabinet:       b.Cabinet,
			Post:          b.Post,
			Education:     b.Education,
			Certificate:   b.Certificate,
			Qualification: b.Qualification,
		}
		pull = append(pull, f)

	}
	c.IndentedJSON(http.StatusOK, pull)
}
func GetPersona(c *gin.Context) {

	key := c.Param("idbranch")
	a := bson.M{"cabinet": key}
	var f = database.PersonalJ{}
	var pull = []database.PersonalJ{}
	var res, _ = database.GetPer(a, database.CollectionPersonal)
	for _, b := range res {
		var stringObjectID = b.ID.Hex()
		f = database.PersonalJ{
			ID:         stringObjectID,
			Surname:    b.Surname,
			Name:       b.Name,
			Patronymic: b.Patronymic,
			Post:       b.Post,
		}
		pull = append(pull, f)

	}
	c.IndentedJSON(http.StatusOK, pull)
}
func Getbranch(c *gin.Context) {

	key := c.Param("branch")
	fmt.Println(key)
	a := bson.M{"branch": key}
	var pull = []database.BranchJ{}
	var f = database.BranchJ{}
	var res, _ = database.Getbra(a, database.CollectionBranch)
	for _, b := range res {
		var stringObjectID = b.ID.Hex()
		f = database.BranchJ{
			ID:       stringObjectID,
			Branch:   b.Branch,
			IdBranch: b.IdBranch,
			Name:     b.Name,
		}
		pull = append(pull, f)
	}
	c.IndentedJSON(http.StatusOK, pull)
}
func Newpersonal(c *gin.Context) {
	var input = database.PersonalJ{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionPersonal)
}
func Newbranch(c *gin.Context) {
	var input = database.BranchJ{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionBranch)
}
func Services(c *gin.Context) {

	key := c.Param("Ser")
	fmt.Println(key)
	a := bson.M{"type": key}
	var res, _ = database.GetSer(a, database.CollectionServices)

	c.IndentedJSON(http.StatusOK, res.Image)
}
func Newcontact(c *gin.Context) {
	var input = database.Contactsj{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionContacts)
}
func NEWServices(c *gin.Context) {

	var input = database.Services{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = primitive.NewObjectID()
	database.CreateTask(input, database.CollectionServices)
}
func Newnews(c *gin.Context) {
	var input = database.Newj{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionNew)
}
func Newprevention(c *gin.Context) {
	var input = database.Newj{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionPrevention)
}
func Reverse(input []*database.New) []*database.New {
	var output []*database.New

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}
func getNews(c *gin.Context) {
	a := bson.M{}
	var res, _ = database.GetNews(a, database.CollectionNew)
	res = Reverse(res)
	var f = database.Newj{}
	var pull = []database.Newj{}
	for _, b := range res {
		var stringObjectID = b.ID.Hex()

		f = database.Newj{
			ID:    stringObjectID,
			Title: b.Title,
			Image: b.Image,
			Text:  b.Text,
			Date:  b.Date,
		}
		pull = append(pull, f)

	}
	c.IndentedJSON(http.StatusOK, pull)
}
func getNew(c *gin.Context) {
	key := c.Param("IDnews")
	objID, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		panic(err)
	}
	a := bson.M{"_id": objID}

	var res, _ = database.GetNews(a, database.CollectionNew)
	var f = database.Itemj{}
	var pull = []database.Itemj{}

	for _, b := range res {

		for _, c := range b.Content {
			f = database.Itemj{
				Type:    c.Type,
				Content: c.Content,
			}
			pull = append(pull, f)
		}

	}
	c.IndentedJSON(http.StatusOK, pull)
}
func getPrevention(c *gin.Context) {
	a := bson.M{}
	var res, _ = database.GetNews(a, database.CollectionPrevention)
	res = Reverse(res)
	var f = database.Newj{}
	var pull = []database.Newj{}
	for _, b := range res {

		var stringObjectID = b.ID.Hex()

		f = database.Newj{
			ID:    stringObjectID,
			Title: b.Title,
			Image: b.Image,
			Text:  b.Text,
			Date:  b.Date,
		}

		pull = append(pull, f)

	}

	c.IndentedJSON(http.StatusOK, pull)
}
func getRead(c *gin.Context) {
	key := c.Param("IDprevention")
	objID, err := primitive.ObjectIDFromHex(key)
	if err != nil {
		panic(err)
	}
	a := bson.M{"_id": objID}

	var res, _ = database.GetNews(a, database.CollectionPrevention)
	var f = database.Itemj{}
	var pull = []database.Itemj{}

	for _, b := range res {

		for _, c := range b.Content {
			f = database.Itemj{
				Type:    c.Type,
				Content: c.Content,
			}
			pull = append(pull, f)

		}

	}
	c.IndentedJSON(http.StatusOK, pull)
}
func getContacts(c *gin.Context) {
	a := bson.M{}
	var res, _ = database.GetContacts(a, database.CollectionContacts)
	var f = database.Contactsj{}
	var pull = []database.Contactsj{}
	for _, b := range res {

		// var stringObjectID = b.ID.Hex()

		f = database.Contactsj{
			// ID:     stringObjectID,
			Name:   b.Name,
			Number: b.Number,
		}
		pull = append(pull, f)

	}

	c.IndentedJSON(http.StatusOK, pull)
}
func getDocuments(c *gin.Context) {
	key := c.Param("Doc")

	a := bson.M{"search": key}

	var res, _ = database.GetDocuments(a, database.CollectionDocuments)

	c.IndentedJSON(http.StatusOK, res)
}
func NewDocuments(c *gin.Context) {
	var input = database.Documents{}
	input.ID = primitive.NewObjectID()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.CreateTask(input, database.CollectionDocuments)
}
