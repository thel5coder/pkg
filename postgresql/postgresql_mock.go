package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"time"
)

type ConnectionMock struct {
	db     *sql.DB
	tx     *sql.Tx
	dbMock sqlmock.Sqlmock
}

func NewConnectionMock() IConnection {
	return &ConnectionMock{}
}

func (c *ConnectionMock) Connect() (IConnection, error) {
	var err error
	c.db, c.dbMock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return c, nil
}

func (c *ConnectionMock) Pool() {
	c.db.SetMaxOpenConns(5)
	c.db.SetMaxIdleConns(3)
	c.db.SetConnMaxLifetime(time.Duration(10) * time.Second)
}

func (c *ConnectionMock) Migration(migrationDirectory string) {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationDirectory,
	}
	n, err := migrate.Exec(c.db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Error migration := ", err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (ConnectionMock) GetDbInstance() *sql.DB {
	panic("implement me")
}

func (c *ConnectionMock) Begin() (err error) {
	c.tx, err = c.db.Begin()

	return err
}

func (c *ConnectionMock) GetTx() *sql.Tx {
	return c.tx
}

func (c *ConnectionMock) Commit() (err error) {
	err = c.tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (c *ConnectionMock) RollBack() (err error) {
	err = c.tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}
