package db

import (
	"context"
	"countries-api/domain"
	"countries-api/entity"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const limit = int64(10)

type Database struct {
	Db mongo.Client
}

func (mdb *Database) init() {
	uri := ""
	if os.Getenv("CONTAINER_MODE") == "true" {
		uri = "mongodb://db:27017/"
	} else if os.Getenv("DATABASE_URL") == "" {
		uri = "mongodb://localhost:27017/"
	} else {
		uri = os.Getenv("DATABASE_URL")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Printf("Database Started on %v", uri)
	mdb.Db = *client
}

func NewDb() domain.DbInterface {
	db := new(Database)
	db.init()
	return db
}

// Create handles the  addition of new countries
func (mdb *Database) Create(country entity.Country) (*entity.Country, error) {
	coll := mdb.Db.Database("country-list").Collection("countries")
	_, err := coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	res, err := coll.InsertOne(context.TODO(), country)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*res)
	return &country, nil
}

func (mdb *Database) Find(id string) (*entity.Country, error) {

	coll := mdb.Db.Database("country-list").Collection("countries")
	filter := bson.D{{"_id", id}}

	var res entity.Country
	err := coll.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (mdb *Database) FindMany(page int) ([]entity.Country, error) {

	l := limit
	skip := int64(page*int(limit) - int(limit))
	option := options.FindOptions{Limit: &l, Skip: &skip}
	coll := mdb.Db.Database("country-list").Collection("countries")

	cursor, err := coll.Find(context.TODO(), bson.D{}, &option)
	if err != nil {
		log.Println("here")
		return nil, err
	}

	var countries []entity.Country

	if err = cursor.All(context.TODO(), &countries); err != nil {
		return nil, errors.New("unable to fetch")
	}

	return countries, nil

}
func (mdb *Database) Update(country entity.Country, id string) (interface{}, error) {

	filter := bson.D{{"_id", id}}
	coll := mdb.Db.Database("country-list").Collection("countries")
	update, err := coll.ReplaceOne(context.TODO(), filter, country)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(update)

	return update, nil
}
