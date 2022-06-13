package migration

import "net/url"

const (
	DefaultMigrationsDir = "./configs/migrations"
)

type Option struct {
	URL                 *url.URL
	MigrationsDir       string
	MigrationsTableName string
	Verbose             bool
	AutoDumpSchema      *bool
	WaitBefore          *bool
}

var (
	tempTrue  = true
	tempFalse = false
)

func (o *Option) Default() *Option {
	if o.MigrationsDir == "" {
		o.MigrationsDir = DefaultMigrationsDir
	}

	if o.AutoDumpSchema == nil {
		o.AutoDumpSchema = &tempTrue
	}

	if o.WaitBefore == nil {
		o.WaitBefore = &tempTrue
	}

	return o
}

type FnOpt func(*Option) (err error)

func WithURL(rawURL string) FnOpt {
	return func(o *Option) (err error) {
		urlInst, err := url.Parse(rawURL)
		if err != nil {
			return
		}
		o.URL = urlInst
		return
	}
}

func WithMigrationsDir(migrationsDir string) FnOpt {
	return func(o *Option) (err error) {
		o.MigrationsDir = migrationsDir
		return
	}
}

func WithMigrationsTableName(migrationTableName string) FnOpt {
	return func(o *Option) (err error) {
		o.MigrationsTableName = migrationTableName
		return
	}
}

func EnableVerbosity() FnOpt {
	return func(o *Option) (err error) {
		o.Verbose = true
		return
	}
}

func DisableAutoDumpSchema() FnOpt {
	return func(o *Option) (err error) {
		o.AutoDumpSchema = &tempFalse
		return
	}
}

func DisableWaitBefore() FnOpt {
	return func(o *Option) (err error) {
		o.WaitBefore = &tempFalse
		return
	}
}
