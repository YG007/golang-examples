package models

type Object interface {
	GetKind() string
	GetID() string
	GetName() string
	SetID(string)
	SetName(string)
}
