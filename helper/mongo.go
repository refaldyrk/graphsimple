package helper

import (
	"context"
	"github.com/qiniu/qmgo"
)

func ConnectMongo() *qmgo.Database {
	cli, err := qmgo.NewClient(context.Background(), &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		panic(err)
	}

	return cli.Database("GraphQL")
}
