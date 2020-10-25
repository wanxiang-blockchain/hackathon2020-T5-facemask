package db

import (
	"context"
	"data-manager/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var DefaultDB *DB = newDB()

func init() {
	optionUri := options.Client().ApplyURI(config.Config.DBConf.Uri())
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	cli, err := mongo.Connect(ctx, optionUri)
	if nil != err {
		panic(err)
	}

	DefaultDB.client = cli
	DefaultDB.Db = cli.Database(config.Config.DBConf.DBName)

	if err = DefaultDB.Ping(); nil != err {
		panic(err)
	}

	//var res model.User
	//ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//res.Name = "admin"
	//res.Password = "admin"
	//collection := DefaultDB.Collection("user")
	//if _, err := collection.InsertOne(ctxBd, res); err != nil {
	//	//fmt.Printf("InsertOne插入的消息ID:%v\n", insertOneRes.InsertedID)
	//	logrus.Errorln(err)
	//
	//}

	logrus.Info("db successfully connected and pinged.")
}

func newDB() *DB {
	return &DB{}
}

type DB struct {
	client *mongo.Client
	Db     *mongo.Database
}

func (this *DB) Collection(name string) *mongo.Collection {
	return this.Db.Collection(name)
}

func (this *DB) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	if err := this.client.Ping(ctx, readpref.Primary()); nil != err {
		return err
	}

	return nil
}

func BuildOptionsByQuery(pageIndex, pageSize int64) *options.FindOptions {
	findOps := options.Find()
	findOps.SetSkip((pageIndex - 1) * pageSize)
	findOps.SetLimit(pageSize)

	return findOps
}
