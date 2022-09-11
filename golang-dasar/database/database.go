package database

var connection string

// func init akan diakses ketika package diakses
func init() {
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}
