package simple_google_wire

/*
Multiple Binding
When doing dependency injection, sometimes there is a case that we create several providers with the same type (struct),
Google Wire doesn't support multiple binding with the same type,
so we need to use alias type for multiple binding.
*/

type Database struct {
	Name string
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	// Convert to DatabaseMongoSQL alias
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	// Convert to DatabasePostgreSQL alias
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

// use alias to multiple binding
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL // alternatively *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB    // alternatively *DatabaseMongoSQL
}

func NewDatabaseRepository(
	postgreSQL *DatabasePostgreSQL,
	mongoSQL *DatabaseMongoDB,
) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongoSQL}
}
