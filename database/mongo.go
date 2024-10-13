package database

import (
	"context"
	"fmt"
	"notes-back/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	connectionString string
	dbName           string
	client           *mongo.Client
}

func NewMongoDatabase(connectionString string, dbName string) *MongoDatabase {
	return &MongoDatabase{
		connectionString: connectionString,
		dbName:           dbName,
	}
}

func (db *MongoDatabase) StringToId(id string) (interface{}, error) {
	return primitive.ObjectIDFromHex(id)
}

func (db *MongoDatabase) Connect() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.connectionString))
	if err != nil {
		return err
	}
	db.client = client
	return nil
}

func (db *MongoDatabase) Disconnect() error {
	return db.client.Disconnect(context.Background())
}

func (db *MongoDatabase) GetUserById(userId string) (types.User, error) {
	coll := db.client.Database(db.dbName).Collection("user")

	var user types.User

	err := coll.FindOne(context.Background(), map[string]string{"id": userId}).Decode(&user)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (db *MongoDatabase) GetUserByEmail(email string) (types.User, error) {
	coll := db.client.Database(db.dbName).Collection("user")

	var user types.User

	err := coll.FindOne(context.Background(), map[string]string{"email": email}).Decode(&user)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

// This function does not implement any validation.
//
// When inserted into the database, sets the UserID field and unsets the Password field
func (db *MongoDatabase) CreateUser(user *types.User) error {
	coll := db.client.Database(db.dbName).Collection("user")

	result, err := coll.InsertOne(context.Background(), user)

	if err != nil {
		if mongoErr, ok := err.(mongo.WriteException); ok {
			for _, we := range mongoErr.WriteErrors {
				if we.Code == 11000 {
					return fmt.Errorf("user with email %s already exists", user.Email)
				}
			}
		} else {
			return err
		}
	}

	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	user.Password = ""

	return err
}

func (db *MongoDatabase) CreateNote(userId string, note *types.Note) (string, error) {

	userIdObj, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return "", err
	}

	note.UserID = userIdObj

	coll := db.client.Database(db.dbName).Collection("note")

	res, err := coll.InsertOne(context.Background(), note)

	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (db *MongoDatabase) UpdateNote(userID string, note *types.Note) error {

	noteIdObj, err := primitive.ObjectIDFromHex(note.ID)

	if err != nil {
		return err
	}
	userIDObj, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	coll := db.client.Database(db.dbName).Collection("note")

	updateFields := bson.D{
		{Key: "title", Value: note.Title},
		{Key: "content", Value: note.Content},
		{Key: "html", Value: note.Html},
		{Key: "preview", Value: note.Preview},
		{Key: "updatedAt", Value: note.UpdatedAt},
	}

	_, err = coll.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: noteIdObj}, {Key: "userID", Value: userIDObj,}}, bson.D{{Key: "$set", Value: updateFields}})

	return err
}

func (db *MongoDatabase) GetUserNotes(userId string) ([]types.Note, error) {
	userIdObj, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	coll := db.client.Database(db.dbName).Collection("note")

	filter := bson.D{{Key: "userID", Value: userIdObj}}
	projection := options.Find().SetProjection(bson.D{{Key: "content", Value: 0}})
	cursor, err := coll.Find(context.Background(), filter, projection)

	if err != nil {
		return nil, err
	}

	var notes []types.Note

	err = cursor.All(context.Background(), &notes)

	return notes, err
}

func (db *MongoDatabase) GetNoteById(noteId string) (types.Note, error) {
	noteIdObj, err := primitive.ObjectIDFromHex(noteId)

	if err != nil {
		return types.Note{}, err
	}

	coll := db.client.Database(db.dbName).Collection("note")

	var note types.Note

	err = coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: noteIdObj}}).Decode(&note)

	return note, err
}
