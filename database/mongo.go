package database

import (
	"context"
	"fmt"
	"notes-back/types"
	"notes-back/types/requestTypes"


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

func (db *MongoDatabase) AddResetCode(email string, code string) error {
	coll := db.client.Database(db.dbName).Collection("resetCode")

	_, err := coll.InsertOne(context.Background(), map[string]string{"email": email, "code": code})

	return err
}

func (db *MongoDatabase) GetResetCode(code string) (string, error) {
	coll := db.client.Database(db.dbName).Collection("resetCode")

	var resetCode map[string]string

	fmt.Println(code)

	err := coll.FindOne(context.Background(), map[string]string{"code": code}).Decode(&resetCode)

	if err != nil {
		return "", err
	}

	return resetCode["email"], nil
}

func (db *MongoDatabase) DeleteResetCode(code string) error {
	coll := db.client.Database(db.dbName).Collection("resetCode")

	_, err := coll.DeleteOne(context.Background(), map[string]string{"code": code})

	return err
}

func (db *MongoDatabase) UpdateUserPassword(email string, password string) error {
	coll := db.client.Database(db.dbName).Collection("user")

	_, err := coll.UpdateOne(context.Background(), 
		bson.D{{Key: "email", Value: email}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: password}}}})
		

	return err
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

func (db *MongoDatabase) UpdateNote(userID string, update *requestTypes.UpdateNote) error {

	noteIdObj, err := primitive.ObjectIDFromHex(update.ID)

	if err != nil {
		return err
	}
	userIDObj, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	coll := db.client.Database(db.dbName).Collection("note")

	updateFields := make(map[string]any)

	GetNoteUpdateFields(update, &updateFields)

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
	projection := options.Find().SetProjection(bson.D{{Key: "content", Value: 0}, {Key: "html", Value: 0}})
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

func (db *MongoDatabase) DeleteNote(ids []string) error {
	var objectIds []primitive.ObjectID

	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err!= nil {
      return err
    }
		objectIds = append(objectIds, objectId)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIds}}

	coll := db.client.Database(db.dbName).Collection("note")
	_, err := coll.DeleteMany(context.Background(), filter)

	return err
}

func (db *MongoDatabase) DeleteNoteById(id string) (error) {
	noteIdObj, err := primitive.ObjectIDFromHex(id)

  if err!= nil {
    return err
  }

  coll := db.client.Database(db.dbName).Collection("note")
  _, err = coll.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: noteIdObj}})

  return err
}

func (db *MongoDatabase) CreateCategory(category *types.Category) error {

	userIdObj, err := primitive.ObjectIDFromHex(category.UserID.(string))

	if err != nil {
		return err
	}

	category.UserID = userIdObj

	coll := db.client.Database(db.dbName).Collection("category")

	res, err := coll.InsertOne(context.Background(), category)

	if err != nil {
		if mongoErr, ok := err.(mongo.WriteException); ok {
			for _, we := range mongoErr.WriteErrors {
				if we.Code == 11000 {
					return fmt.Errorf("category with name %s already exists", category.Name)
				}
			}
		} else {
			return err
		}
	}

	category.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (db *MongoDatabase) GetCategories(userId string) ([]types.Category, error) {
	userIdObj, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	coll := db.client.Database(db.dbName).Collection("category")

	filter := bson.D{{Key: "userID", Value: userIdObj}}
	cursor, err := coll.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var categories []types.Category

	err = cursor.All(context.Background(), &categories)

	fmt.Println(categories)

	return categories, err
}

func (db *MongoDatabase) DeleteCategory(id string) error {
	categoryIdObj, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	coll := db.client.Database(db.dbName).Collection("category")

	_, err = coll.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: categoryIdObj}})

	return err
}