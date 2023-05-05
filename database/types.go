package database

import (
	"database/sql"
	"embed"
	"math/big"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
)

var content embed.FS

type DatabaseDriver string

// Database Parameters Needed To Setup The Network
type DatabaseOpts struct {
	MaxConnections        int
	MaxIdleConnections    int
	MaxIdleTimeConnection time.Duration
}

type DatabaseInterface struct {
	DB     *sql.DB // Function so we have functions on top of it
	Opts   DatabaseOpts
	Driver DatabaseDriver
	Log    logrus.Entry
	URL    string
}

func (database *DatabaseInterface) NewDatabaseOpts() {

	database.DB.SetMaxOpenConns(database.Opts.MaxConnections)

	database.DB.SetMaxIdleConns(database.Opts.MaxIdleConnections)

	database.DB.SetConnMaxIdleTime(database.Opts.MaxIdleTimeConnection)
}

func (database *DatabaseInterface) DBMigrate() error {
	migrationOpts, err := iofs.New(content, "migrations/")
	if err != nil {
		database.Log.Fatal(err)
		return err
	}

	migration, err := migrate.NewWithSourceInstance("iofs", migrationOpts, database.URL)
	if err != nil {
		database.Log.Fatal(err)
		return err
	}

	defer migration.Close()

	err = migration.Up()
	if err != nil {
		database.Log.Fatal("Database Migrate Error")
		return err
	}

	database.Log.Info("Database Migrate Succesful")
	return nil
}

type ValidatorDeliveredPayloadDatabase struct {
	Slot           uint64
	ProposerPubkey string
	BlockHash      string
	Payload        ExecutionPayload
}

type ValidatorReturnedBlockDatabase struct {
	Signature      string
	Slot           uint64
	BlockHash      string
	ProposerPubkey string
}

type ValidatorDeliveredHeaderDatabase struct {
	Slot           uint64
	Value          big.Int
	BlockHash      string
	ProposerPubkey string
}

type BuilderBlockDatabase struct {
	Slot             uint64
	BuilderPubkey    string
	BuilderSignature string
	RPBS             RPBS
	TransactionByte  string
}
