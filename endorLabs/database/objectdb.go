package database

import (
	"context"
	"endorLabs/models"
)

type ObjectDB interface {
	Store(ctx context.Context, object models.Object) error
	GetObjectByID(ctx context.Context, id string) (models.Object, error)
	GetObjectByName(ctx context.Context, name string) (models.Object, error)
	ListObjects(ctx context.Context, kind string) ([]models.Object, error)
	DeleteObject(ctx context.Context, id string) error
}
