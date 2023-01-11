package database

// Repository uses for implement storage of models.

// IRepository - repository interface.
type IRepository interface {
	ModelGenerator() ModelGenerator
}

// Repository - abstract repository struct.
type Repository struct {
	Database *Database
}

// NewQuery - instantiate new query.
func (r *Repository) NewQuery(cr IRepository) Query {
	return r.Database.NewQuery(cr)
}
