package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	VisitCollection *mongo.Collection
	Ctx             = context.TODO()
	dbConn          *gorm.DB
)

func SetupMongoDBConnection() {
	logger := log.NewLogfmtLogger(os.Stderr)

	clientOptions := options.Client().ApplyURI(fmt.Sprint(viper.Get("MONGODB_URI")))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		level.Error(logger).Log("msg", err)

	}
	err = client.Ping(Ctx, nil)
	if err != nil {
		level.Error(logger).Log("msg", err)
	}
	db := client.Database("hello")
	VisitCollection = db.Collection("visit")
}

func GetPostgresqlDBConnection() error {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		fmt.Sprint(viper.Get("PGHOST")),
		fmt.Sprint(viper.Get("PGUSER")),
		fmt.Sprint(viper.Get("PGPASSWORD")),
		fmt.Sprint(viper.Get("PGDBNAME")),
		fmt.Sprint(viper.Get("PGPORT")),
		fmt.Sprint(viper.Get("PGSSLMODE")),
		fmt.Sprint(viper.Get("PGTIMEZONE")))
	pdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := pdb.DB()
	if err != nil {
		return err
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = pdb
	return nil
}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}
