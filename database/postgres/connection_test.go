package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "test",
		DBName:   "test",
		SSLMode:  "disable",
	}

	t.Run("Valid Connection", func(t *testing.T) {
		conn, err := NewConnection(config)
		if err != nil {
			t.Skip("Skipping test: Could not connect to database")
		}
		defer conn.Close()

		// Test connection by pinging
		err = conn.DB().Ping()
		assert.NoError(t, err, "Should connect to database successfully")
	})

	t.Run("Invalid Connection", func(t *testing.T) {
		config.Host = "invalid-host"
		_, err := NewConnection(config)
		assert.Error(t, err)
	})
}

func TestTransaction(t *testing.T) {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "test",
		DBName:   "test",
		SSLMode:  "disable",
	}

	t.Run("Successful Transaction", func(t *testing.T) {
		conn, err := NewConnection(config)
		if err != nil {
			t.Skip("Skipping test: Could not connect to database")
		}
		defer conn.Close()

		err = conn.Transaction(func(tx *sql.Tx) error {
			// Create a test table
			_, err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS test_table (
					id SERIAL PRIMARY KEY,
					name TEXT
				)
			`)
			return err
		})
		assert.NoError(t, err, "Transaction should complete successfully")
	})

	t.Run("Transaction Rollback", func(t *testing.T) {
		conn, err := NewConnection(config)
		if err != nil {
			t.Skip("Skipping test: Could not connect to database")
		}
		defer conn.Close()

		err = conn.Transaction(func(tx *sql.Tx) error {
			return errors.New("force rollback")
		})
		assert.Error(t, err)
	})
}
