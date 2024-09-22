package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

/*
  - MigrateBooks runs the database migration for the Books model.

*   It uses GORM's AutoMigrate function to automatically migrate the schema.
*   If the table does not exist, AutoMigrate will:
*     - Create a new table based on the Books struct definition.
*     - Add all necessary fields, column types, indexes, and constraints.
*
*   If the table already exists, AutoMigrate will:
*     - Add any new columns that are defined in the Books struct but not in the table.
*     - Modify column types to match the struct (depending on database support).
*     - Add missing indexes and constraints.
*
*   Note: AutoMigrate will NOT remove existing columns or indexes that are no longer present in the Books struct.
*   If the migration fails, it returns an error. Otherwise, it returns nil.
*
* Parameters:
*   db *gorm.DB - The GORM database connection instance.
*
* Returns:
*   error - Returns an error if the migration fails, otherwise returns nil.
*/
func MigrateBooks(db *gorm.DB) error {
	// Run the auto-migration for the Books struct, creating the table if it doesn't exist
	err := db.AutoMigrate(&Books{})
	// Return the error if AutoMigrate fails, or nil if it's successful
	return err
}
