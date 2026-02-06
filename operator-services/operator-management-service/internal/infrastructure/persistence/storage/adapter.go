package storage

import (
	"context"
	"fmt"
	"time"
	operatormodel "ylanzinhoy-operator-management/internal/core/model/operatorModel"
	operatorEntity "ylanzinhoy-operator-management/internal/domain/entity/operator"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Adapter struct {
	dbName   string
	collName string
	coll     *mongo.Collection
}

func NewAdapter(collName, dbName string, coll *mongo.Collection) *Adapter {
	return &Adapter{
		collName: collName,
		dbName:   dbName,
		coll:     coll,
	}
}

func (a *Adapter) Insert(ctx context.Context, entity *operatorEntity.Operator) (any, error) {

	operatorModel := operatormodel.EntityToModel(entity)
	now := time.Now().UTC()
	operatorModel.CreatedAt = now
	operatorModel.UpdatedAt = now

	insertResult, err := a.coll.InsertOne(ctx, operatorModel)

	if err != nil {
		return "", fmt.Errorf("falha ao inserir operador no banco de dados")
	}

	if insertResult.InsertedID == nil {
		return "", mongo.ErrNilDocument
	}

	return insertResult.InsertedID, nil

}

func (a *Adapter) FindByID(ctx context.Context, id string) (*operatorEntity.Operator, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("id invalido: %w", err)
	}

	var operatorModel operatormodel.OperatorModel
	if err := a.coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(&operatorModel); err != nil {
		return nil, err
	}

	entity, err := operatormodel.ModelToEntity(operatorModel)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (a *Adapter) FindAll(ctx context.Context) ([]*operatorEntity.Operator, error) {
	cursor, err := a.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var models []operatormodel.OperatorModel
	if err := cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	entities := make([]*operatorEntity.Operator, 0, len(models))
	for _, model := range models {
		entity, err := operatormodel.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (a *Adapter) UpdateByID(ctx context.Context, id string, entity *operatorEntity.Operator) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("id invalido: %w", err)
	}

	operatorModel := operatormodel.EntityToModel(entity)
	operatorModel.UpdatedAt = time.Now().UTC()

	updateResult, err := a.coll.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": operatorModel})
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (a *Adapter) DeleteByID(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("id invalido: %w", err)
	}

	deleteResult, err := a.coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (a *Adapter) UpdatePassword(ctx context.Context, id string, entity *operatorEntity.Operator) error {

	now := time.Now().UTC()
	operatorModel := operatormodel.EntityToModel(entity)
	operatorModel.UpdatedAt = now

	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("id invalido: %w", err)
	}

	update := bson.M{
		"$set": bson.M{"password": entity.Password.Value},
	}

	result, err := a.coll.UpdateOne(ctx, bson.M{"_id": objectId}, update)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil

}
