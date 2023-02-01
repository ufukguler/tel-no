package database

import (
	"context"
	. "github.com/gobeam/mongo-go-pagination"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"telno/config"
)

var Database DB

type DB struct {
	client   *mongo.Client
	database *mongo.Database
	ctx      context.Context
}

func (db *DB) Connect() {
	connectionString := config.GetMongoConnectionString()
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panic(err)
	}

	dbName := config.GetEnv("MONGODB_DATABASE")
	Database.client = client
	Database.ctx = context.TODO()
	Database.database = client.Database(dbName)

	err = client.Ping(Database.ctx, nil)
	if err != nil {
		log.Panic(err)
	}
	log.Info("[DB] Connected. DB_NAME: ", dbName)
}

func (db *DB) Truncate(collection string) error {
	col := db.database.Collection(collection)
	_, err := col.DeleteMany(db.ctx, bson.M{})
	return err
}

func (db *DB) InsertOne(collection string, object interface{}) error {
	col := db.database.Collection(collection)
	_, err := col.InsertOne(db.ctx, object)
	return err
}

func (db *DB) FindOne(collection string, filter interface{}, object interface{}) error {
	col := db.database.Collection(collection)
	err := col.FindOne(db.ctx, filter).Decode(object)
	return err
}

func (db *DB) Find(collection string, filter interface{}, object interface{}) error {
	col := db.database.Collection(collection)
	find, err := col.Find(db.ctx, filter)
	if err != nil {
		return err
	}
	err = find.All(db.ctx, object)
	return err
}

func (db *DB) FindByOptions(collection string, filter interface{}, object interface{}, findOptions options.FindOptions) error {
	col := db.database.Collection(collection)
	find, err := col.Find(db.ctx, filter, &findOptions)
	if err != nil {
		return err
	}
	err = find.All(db.ctx, object)
	return err
}

func (db *DB) IsDocExist(collection string, filter interface{}) (bool, error) {
	col := db.database.Collection(collection)
	count, err := col.CountDocuments(db.ctx, filter)
	if err != nil {
		return false, err
	}
	if count >= 1 {
		return true, nil
	}
	return false, nil
}

func (db *DB) AggregateQuery(collection string, filter interface{}, object interface{}) error {
	col := db.database.Collection(collection)
	aggregate, err := col.Aggregate(db.ctx, filter)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return aggregate.All(db.ctx, object)
}

func (db *DB) UpdateOne(collection string, filter interface{}, updateStatement interface{}) error {
	col := db.database.Collection(collection)
	_, err := col.UpdateOne(db.ctx, filter, updateStatement)
	return err
}

func (db *DB) CountDocuments(collection string, filter interface{}) (int64, error) {
	col := db.database.Collection(collection)
	return col.CountDocuments(db.ctx, filter)
}

func (db *DB) DeleteMany(collection string, filter interface{}) (int64, error) {
	col := db.database.Collection(collection)
	result, err := col.DeleteMany(db.ctx, filter)
	return result.DeletedCount, err
}

func (db *DB) DeleteOne(collection string, filter interface{}) error {
	col := db.database.Collection(collection)
	_, err := col.DeleteOne(db.ctx, filter)
	return err
}

func (db *DB) FindByPageable(collection string, limit int64, page int64) (*PaginatedData, error) {
	col := db.database.Collection(collection)
	return New(col).Limit(limit).Page(page).Aggregate()
}

func (db *DB) FindByPageableMatch(collection string, limit int64, page int64, filter interface{}, filter2 interface{}) (*PaginatedData, error) {
	col := db.database.Collection(collection)
	return New(col).Limit(limit).Page(page).Aggregate(filter, filter2)
}
