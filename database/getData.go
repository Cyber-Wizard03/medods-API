package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// type GetData interface{

// 	cur, err := collection.Find(ctx, filter)

// }

func GetInfo(filter interface{}, collection *mongo.Collection) (Info, error) {

	data := Info{}
	err := collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	// This error means your query did not match any documents.
		// 	return session{}
		// }
		panic(err)
	}
	return data, err
}

func GetVacancy(filter interface{}, collection *mongo.Collection) ([]*Vacancy, error) {
	var documents = []*Vacancy{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := Vacancy{}
		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}
func GetFeedback(filter interface{}, collection *mongo.Collection) ([]*MesFeedback, error) {
	var documents = []*MesFeedback{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := MesFeedback{}
		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}
func GetPer(filter interface{}, collection *mongo.Collection) ([]*Personal, error) {
	var documents = []*Personal{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := Personal{}
		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}
func GetSer(filter interface{}, collection *mongo.Collection) (Services, error) {
	data := Services{}
	err := collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	// This error means your query did not match any documents.
		// 	return session{}
		// }
		panic(err)
	}
	return data, err
}

func Getbra(filter interface{}, collection *mongo.Collection) ([]*Branch, error) {

	var documents = []*Branch{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := Branch{}
		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}

func GetNews(filter interface{}, collection *mongo.Collection) ([]*New, error) {

	var documents = []*New{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := New{}
		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}

func GetContacts(filter interface{}, collection *mongo.Collection) ([]*Contacts, error) {

	var documents = []*Contacts{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := Contacts{}

		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}

func GetDocuments(filter interface{}, collection *mongo.Collection) ([]*Documents, error) {

	var documents = []*Documents{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return documents, err
	}

	for cur.Next(ctx) {
		document := Documents{}

		err := cur.Decode(&document)

		if err != nil {
			return documents, err
		}

		documents = append(documents, &document)
	}

	if err := cur.Err(); err != nil {
		return documents, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(documents) == 0 {
		return documents, mongo.ErrNoDocuments
	}

	return documents, nil
}
