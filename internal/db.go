package internal

import (
	"log"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepository interface {
	InsertCourse(c Course) (string, error)
	GetCourses() ([]Course, error)
	GetCourse(id string) (Course, error)
	DeleteCourse(id string) error
}

type MongoService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (self *MongoService) CreateCourse() int {
	return 0
}

func NewDB(dbUri string) CourseRepository {
	// Create client and connect to db server
	clientOpts := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database(os.Getenv("MONGO_DB")).Collection("users")
	return &MongoService{client: client, collection: coll}	
}

func (self *MongoService) InsertCourse(c Course) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := self.collection.InsertOne(ctx, bson.M{"title": c.Title, "description": c.Description})
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (self *MongoService) GetCourses() ([]Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := self.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var courses []Course
	for cursor.Next(ctx) {
		var c Course
		cursor.Decode(&c)
		courses = append(courses, c)
	}
	return courses, nil
}

func (self *MongoService) GetCourse(id string) (Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Course{}, err
	}
	var course Course
	err = self.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&course)
	return course, err
}

func (self *MongoService) DeleteCourse(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = self.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
