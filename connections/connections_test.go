package connections

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

// TestGetMongoConnection valida que se reciba una instancia de *mongo.Database y que se pueda usar
func TestGetMongoConnection(t *testing.T) {
	db := GetMongoConnection(context.WithTimeout(context.Background(), time.Duration(10)*time.Second))
	res, e := db.Collection("credit_details").InsertOne(context.TODO(), primitive.D{
		primitive.E{
			Key:   "text",
			Value: "Hello, everybody! :D",
		},
	})

	if e != nil {
		t.Error(e)
	}
	t.Log(res)
	e = db.Collection("credit_details").Drop(context.TODO())
	if e != nil {
		t.Error(e)
	}

}

// TestGetMongoConnectionBadUri valida que se envíe error cuando no se pudo realizar la conexión
func TestGetMongoConnectionBadUri(t *testing.T) {
	e := os.Setenv("MONGO_URI", "bad_mongo_uri")
	if e != nil {
		t.Error(e)
	}

	defer func() {
		assert.NotNil(t,recover(), "Debió lanzar un error al iniciar la base de datos")
	}()
	db := GetMongoConnection(context.WithTimeout(context.Background(), time.Duration(10)*time.Second))
	assert.Nil(t, db, "No debió inicializar la base de datos")

	e = os.Unsetenv("MONGO_URI")
	if e != nil {
		t.Error(e)
	}
}

// TestGetMongoConnectionPingError valida que se envíe error cuando no se pudo realizar el ping a mongo
func TestGetMongoConnectionPingError(t *testing.T) {
	e := os.Setenv("MONGO_URI", "mongodb://localhost:27018/")
	if e != nil {
		t.Error(e)
	}
	defer func() {
		assert.NotNil(t,recover(), "Debió lanzar un error al iniciar la base de datos")
	}()

	db := GetMongoConnection(context.WithTimeout(context.Background(), 4*time.Second))
	assert.Nil(t, db, "No debió inicializar la base de datos")

	e = os.Setenv("MONGO_URI", "bad_mongo_uri")
	if e != nil {
		t.Error(e)
	}

}