package connections

import (
	"context"
	"github.com/manicar2093/YoFioExamen/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection(ctx context.Context, cancelFunc context.CancelFunc) *mongo.Database {
	defer cancelFunc()
	opt := options.Client().ApplyURI(utils.GetEnvVar("MONGO_URI", "mongodb://localhost:27017/"))
	conn, e := mongo.Connect(ctx, opt)
	if e != nil {
		utils.LogError.Printf("Error al conectarse a mongodb.\n\tDetalles: %v", e)
		panic(e)
	}

	e = conn.Ping(ctx, nil)
	if e != nil {
		utils.LogError.Printf("Error al conectarse a mongodb.\n\tDetalles: %v", e)
		panic(e)
	}
	return conn.Database("yoFio_prueba_tecnica")

}
