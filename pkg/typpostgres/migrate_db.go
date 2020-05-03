package typpostgres

import (
	"github.com/golang-migrate/migrate"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func cmdMigrateDB(c *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:   "migrate",
		Usage:  "Migrate Database",
		Action: c.ActionFunc("PG", migrateDB),
	}
}

func migrateDB(c *typbuildtool.CliContext) (err error) {
	var (
		migration *migrate.Migrate
		cfg       *Config
	)

	if cfg, err = retrieveConfig(); err != nil {
		return
	}

	sourceURL := "file://" + DefaultMigrationSource
	c.Infof("Migrate database from source '%s'", sourceURL)
	if migration, err = migrate.New(sourceURL, dataSource(cfg)); err != nil {
		return err
	}
	defer migration.Close()
	return migration.Up()
}