package storage

import (
	"context"
	"log"
	"time"

	"auditlog/config"
	"auditlog/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

// The function establishes a MongoDB connection and sets the collection for audit log events.
func init() {
	// Establish MongoDB connection
	if config.IsClustered {

		clientOptions := options.Client().ApplyURI(config.MongodbURI)

		var err error
		client, err = mongo.NewClient(clientOptions)

		if err != nil {
			log.Fatalf("Storage: Failed to create MongoDB client config.MongodbURI: %v err: %v", config.MongodbURI, err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Fatal("Storage: Failed to connect to MongoDB:", err)
		}

		// Set the collection
		database := client.Database("auditlog_db")
		collection = database.Collection("auditlog_event_col")
		log.Printf("Storage: Connected to Mongodb and created database and collection db:%v, coll:%v\n", database, collection)
	}
}

// The function inserts a document into a MongoDB collection.
func InsertDoc(doc models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	raw, err := bson.Marshal(doc)
	if err != nil {
		log.Println("Storage: Failed to marshal document:", err)
		return err
	}

	_, err = collection.InsertOne(ctx, raw)
	if err != nil {
		log.Println("Storage: Failed to insert document:", err)
		return err
	}

	log.Printf("Storage: Document inserted successfully: %v\n", doc)
	return nil
}

// The function retrieves documents from a MongoDB collection based on a given filter and returns them
// as a slice of models.Event structs.
func FindDoc(filter map[string]interface{}) ([]models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Storage: Received request to retrieve document with filter: %v\n", filter)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Storage: Failed to find documents:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Event
	for cursor.Next(ctx) {
		var result models.Event
		err := cursor.Decode(&result)
		if err != nil {
			log.Println("Storage: Failed to decode document:", err)
			return nil, err
		}
		results = append(results, result)
	}

	log.Printf("Storage: Documents found: %v\nTotal documents: %d\n", results, len(results))
	return results, nil
}

// The function deletes a document from a collection in a MongoDB database based on a given filter.
// TODO: not used yet
func DeleteDoc(filter map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Storage: Failed to delete document:", err)
		return err
	}

	log.Println("Storage: Document deleted successfully")
	return nil
}

// The function updates a document in a collection using a filter and an update object.
// TODO: not used yet
func UpdateDoc(filter map[string]interface{}, update models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Storage: Failed to update document:", err)
		return err
	}

	log.Println("Storage: Document updated successfully")
	return nil
}
