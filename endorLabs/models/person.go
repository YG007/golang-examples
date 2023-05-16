package models

import (
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	LastName  string    `json:"last_name"`
	Birthday  string    `json:"birthday"`
	BirthDate time.Time `json:"birth_date"`
}

func (p *Person) GetKind() string {
	return reflect.TypeOf(p).String()
}

func (p *Person) GetID() string {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) SetID(id string) {
	p.ID = id
}

func (p *Person) SetName(name string) {
	p.Name = name
}
