package projects

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (project *Project) AddToLocal() {
	if project.Exists() {
		project.UpdateLocal()
		return
	}
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Projects")

	_, err := collection.InsertOne(ctx, project)
	variables.SaveError(err, "projects", "project.AddToLocal")
}


func (project *Project) UpdateLocal() {
	ctx := context.Background()
	collection := variables.Local.Database("Interphlix").Collection("Projects")


	filter := bson.M{"_id": bson.M{"$eq": project.ID}}
	update := bson.M{"$set": project}
	_, err := collection.UpdateOne(ctx, filter, update)
	variables.SaveError(err, "projects", "project.UpdateLocal")
}