package store

import "fmt"

type ErrInvalidInput struct {
	Entity string
	Field  string
	Value  interface{}
}

func NewErrInvalidInput(entity, field string, value interface{}) *ErrInvalidInput {
	return &ErrInvalidInput{
		Entity: entity,
		Field:  field,
		Value:  value,
	}
}

func (e *ErrInvalidInput) Error() string {
	return fmt.Sprintf("invalid inpit: entity: %s field: %s value: %s", e.Entity, e.Field, e.Value)
}
