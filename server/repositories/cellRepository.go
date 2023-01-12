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

// CastCell - implements model handler to cast cell model.
func (r *CellRepository) CastCell(h CellHandler) database.ModelHandler {
	return func(m database.IModel) {
		c := m.(*models.Cell)

		// destroy model if model is empty.
		if c.ID == 0 {
			c = nil
		}

		h(c)
	}
}

// FindAll - find all models in repository.
func (r *CellRepository) FindAll() *database.Collection {
	q := r.NewQuery(r)

	q.BuildSelect()

	return q.SelectAll()
}

// FindOneByColumnRow - find one model by column and
// row in repository.
func (r *CellRepository) FindOneByColumnRow(column *string, row *int) *database.Item {
	fq := r.NewQuery(r)

	fq.BuildSelect()
	fq.Where(database.WhereCondition{
		Conditions: []string{"`column` = ?", "`row` = ?"},
		Op:         "AND",
	})

	return fq.Select(column, row)
}

// Save - save model in repository.
func (r *CellRepository) Save(c *models.Cell) {
	i := r.FindOneByColumnRow(&c.Column, &c.Row)

	var fc *models.Cell
	i.One(r.CastCell(func(c *models.Cell) {
		fc = c
	}))

	// Chose or insert or update model.
	q := r.NewQuery(r)
	if fc == nil {
		q.BuildInsert()

		c.ID = int(q.Insert(c))
	} else {
		q.BuildUpdate()
		q.Where(database.WhereCondition{
			Conditions: []string{"`id` = ?"},
		})

		q.Update(&fc.ID, c)
	}
}

// NewCellRepository - instantiate new cell repository.
func NewCellRepository() *CellRepository {
	return &CellRepository{
		Repository: database.Repository{
			Database: database.Db,
		},
	}
}
