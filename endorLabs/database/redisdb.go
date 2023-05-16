package database

import (
	"context"
	"endorLabs/models"
	"fmt"
	"reflect"

	"github.com/go-redis/redis/v8"
)

type RedisObjectDB struct {
	client *redis.Client
}

func NewRedisObjectDB(client *redis.Client) *RedisObjectDB {
	return &RedisObjectDB{
		client: client,
	}
}

func (r *RedisObjectDB) Store(ctx context.Context, object models.Object) error {
	id := object.GetID()
	return r.client.HMSet(ctx, id, map[string]interface{}{
		"name": object.GetName(),
		"kind": reflect.TypeOf(object).Elem().Name(),
	}).Err()
}

func (r *RedisObjectDB) GetObjectByID(ctx context.Context, id string) (models.Object, error) {
	result, err := r.client.HGetAll(ctx, id).Result()
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("object not found")
	}
	objectKind := result["kind"]
	objectName := result["name"]

	switch objectKind {
	case "Person":
		person := &models.Person{
			ID:   id,
			Name: objectName,
		}
		return person, nil
	case "Animal":
		animal := &models.Animal{
			ID:   id,
			Name: objectName,
		}
		return animal, nil
	default:
		return nil, fmt.Errorf("unknown object kind: %s", objectKind)
	}
}

func (r *RedisObjectDB) GetObjectByName(ctx context.Context, name string) (models.Object, error) {
	ids, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		result, err := r.client.HGetAll(ctx, id).Result()
		if err != nil {
			continue
		}
		objectName := result["name"]
		if objectName == name {
			return r.GetObjectByID(ctx, id)
		}
	}
	return nil, fmt.Errorf("object not found")
}

func (r *RedisObjectDB) ListObjects(ctx context.Context, kind string) ([]models.Object, error) {
	ids, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	var objects []models.Object
	for _, id := range ids {
		result, err := r.client.HGetAll(ctx, id).Result()
		if err != nil {
			continue
		}
		objectKind := result["kind"]
		if objectKind == kind {
			var object models.Object
			switch objectKind {
			case "Person":
				object = &models.Person{
					ID:   id,
					Name: result["name"],
				}
			case "Animal":
				object = &models.Animal{
					ID:   id,
					Name: result["name"],
				}
			default:
				return nil, fmt.Errorf("unknown object kind: %s", objectKind)
			}
			objects = append(objects, object)
		}
	}
	return objects, nil
}

func (r *RedisObjectDB) DeleteObject(ctx context.Context, id string) error {
	return r.client.Del(ctx, id).Err()
}
