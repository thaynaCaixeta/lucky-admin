package database

import (
	"embed"

	migrate "github.com/rubenv/sql-migrate"
)

// FS contains the migration files
//
//go:embed *.sql
var dbMigrations embed.FS

// GetMigrationSource returns the migration source.
func GetMigrationSource() *migrate.EmbedFileSystemMigrationSource {
	return &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       ".",
	}
}
