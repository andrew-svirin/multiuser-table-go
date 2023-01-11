package repositories

import (
	"github.com/andrew-svirin/multiuser-table-go/server/models"
	"github.com/andrew-svirin/multiuser-table-go/server/services/database"
)

// CellHandler - type for cell handler.
type CellHandler func(*models.Cell)

// CellRepository - cell repository struct.
type CellRepository struct {
	database.Repository
}

// ModelGenerator - implements repository model generator.
func (r *CellRepository) ModelGenerator() database.ModelGenerator {
	return func() database.IModel {
		return models.NewCell()
	}
}

// CellHandler - implements model handler to cast cell model.
func (r *CellRepository) CellHandler(h CellHandler) database.ModelHandler {
	return func(m database.IModel) {
		c := m.(*models.Cell)

		h(c)
	}
}

// FindAll - find all models in repository.
func (r *CellRepository) FindAll() *database.Collection {
	q := r.NewQuery(r)

	q.BuildSelect()

	return q.Exec()
}

// NewCellRepository - instantiate new cell repository.
func NewCellRepository() *CellRepository {
	return &CellRepository{
		Repository: database.Repository{
			Database: database.Db,
		},
	}
}
