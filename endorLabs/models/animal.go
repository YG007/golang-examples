package models

import (
	"reflect"

	"github.com/google/uuid"
)

type Animal struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	OwnerID string `json:"owner_id"`
}

func (a *Animal) GetKind() string {
	return reflect.TypeOf(a).String()
}

func (a *Animal) GetID() string {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return a.ID
}

func (a *Animal) GetName() string {
	return a.Name
}

func (a *Animal) SetID(id string) {
	a.ID = id
}

func (a *Animal) SetName(name string) {
	a.Name = name
}
