package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Row struct {
	ID       string
	Value    float64
	Currency string
	Date     time.Time
	Tags     []string
}

func NewRow(val float64, tags ...string) *Row {
	return &Row{
		ID:       uuid.New().String(),
		Value:    val,
		Currency: "CAD",
		Date:     time.Now(),
		Tags:     tags,
	}
}

// Will save the information in the DB
func (r *Row) save() {
	_, err := rowCollection.UpdateOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: r.ID}},
		bson.D{{Key: "$set", Value: r}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		fmt.Println("Couldn't save this row:" + err.Error())
		return
	}
}
