package storage

import (
	"fmt" // Used for formatting the connection string (DSN)

	"gorm.io/driver/postgres" // PostgreSQL-specific driver for GORM
	"gorm.io/gorm"            // GORM core library
)

// Config struct holds the configuration needed to connect to the PostgreSQL database
type Config struct {
	Host     string // Database server address
	Port     string // Database server port
	Password string // Password for the database user
	User     string // Username for the database connection
	DBName   string // Name of the PostgreSQL database
	SSLMode  string // SSL mode for the connection (e.g., "disable", "require", etc.)
}

// NewConnection function creates a new connection to the PostgreSQL database using GORM.
// It takes a pointer to a Config struct as input and returns a GORM DB connection or an error.
func NewConnection(config *Config) (*gorm.DB, error) {
	// Build the Data Source Name (DSN) string using the provided configuration.
	// This DSN is required to establish the database connection.
	dsn := fmt.Sprintf("host=%s port=%s password=%s user=%s dbname=%s sslMode=%s",
		config.Host,     // Host of the database server
		config.Port,     // Port of the database server
		config.Password, // Password for the database user
		config.User,     // Username for the database
		config.DBName,   // Name of the database to connect to
		config.SSLMode)  // SSL mode for the connection (e.g., "disable", "require")

	/*
	 * gorm.Open
	 * Purpose: Initializes a new GORM DB instance with the given driver and configuration.
	 * Signature:
	 *     func Open(dialector Dialector, opts ...Option) (*DB, error)
	 *
	 * Driver Function: This is specific to each database system and is used to initialize
	 * the connection with gorm.Open. For example, postgres.Open, mysql.Open, and sqlite.Open
	 * are used for PostgreSQL, MySQL, and SQLite, respectively.
	 *
	 *
	 * postgres.Open as the Driver:
	 * The postgres package provides the PostgreSQL driver for GORM.
	 *
	 * Driver Definition:
	 * The postgres.Open function is used to create a new instance of the PostgreSQL driver
	 * configured with a given Data Source Name (DSN).
	 */

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// If an error occurred while trying to connect, return the error and nil for the DB connection.
	if err != nil {
		return nil, err
	}

	// If the connection is successful, return the GORM DB connection and nil for the error.
	return db, nil
}

/*
Implicit Nil Handling: In Go, if the return type of a function is error, and you return nil, it's understood that no error occurred.
The error type is a special case in Go where returning nil indicates success, and any non-nil value indicates an error.
*/
func CloseConnection(db *gorm.DB) error {
	// Retrieve the underlying *sql.DB instance from GORM
	sqlDB, err := db.DB()
	if err != nil {
		return err // Return error if there's an issue retrieving *sql.DB
	}
	// Return the result of sqlDB.Close()
	return sqlDB.Close() // This returns either nil (success) or an error (failure)
}
