package migrations

import (
	"L0/db/postgres"
)

func DataMigrate() bool {
	row, err := postgres.CheckIfExist()
}
