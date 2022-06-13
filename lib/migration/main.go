package migration

import (
	"github.com/amacneil/dbmate/pkg/dbmate"
)

func New(opts ...FnOpt) (db *dbmate.DB, err error) {
	option := new(Option)

	for _, opt := range opts {
		err = opt(option)
		if err != nil {
			return
		}
	}

	option.Default()

	db = dbmate.New(option.URL)
	db.MigrationsDir = option.MigrationsDir
	db.MigrationsTableName = option.MigrationsTableName
	db.Verbose = option.Verbose
	db.AutoDumpSchema = *option.AutoDumpSchema
	db.WaitBefore = *option.WaitBefore

	return
}
