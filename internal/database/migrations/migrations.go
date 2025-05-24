package migrations

import (
	"embed"

	migrate "github.com/rubenv/sql-migrate"
)

// FS contains the migration files
var FS embed.FS

// GetMigrationSource returns the migration source.
func GetMigrationSource() *migrate.EmbedFileSystemMigrationSource {
	return &migrate.EmbedFileSystemMigrationSource{
		FileSystem: FS,
		Root:       ".",
	}
}
