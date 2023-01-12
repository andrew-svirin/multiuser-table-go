package database

// Model uses to represent custom models.

// IModel - database model interface.
type IModel interface {
	Schema() *Schema
}

// ModelGenerator - model generator type.
type ModelGenerator func() IModel

// ModelHandler - model handler type.
type ModelHandler func(IModel)
