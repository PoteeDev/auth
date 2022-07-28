package storage

import (
	"context"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/entities/models"
	"gopkg.in/mgo.v2/bson"
)

func GetAuthTeam(login string) (models.Entity, error) {

	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var entity models.Entity
	err := col.FindOne(ctx, bson.M{"login": login}).Decode(&entity)
	return entity, err
}
