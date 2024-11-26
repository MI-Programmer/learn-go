package database

var database string

func init() {
	database = "MySQL"
}

func GetDatabase() string {
	return database
}
