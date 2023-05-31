package database

import (
	"crypto/sha256"
	"database/sql"
	"embed"
	"fmt"
	"math/big"
	"time"

	relayTypes "github.com/bsn-eng/pon-golang-types/relay"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/sirupsen/logrus"
)

var Content embed.FS

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
	migrationOpts, err := iofs.New(Content, "migrations/")
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
	Payload        []byte
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
	RPBS             relayTypes.EncodedRPBSSignature
	TransactionByte  string
	Value            uint64
}

func (builderSubmission *BuilderBlockDatabase) Hash() string {
	BuilderBid := fmt.Sprintf("%d,%s,%s,%d",
		builderSubmission.Slot,
		builderSubmission.BuilderPubkey,
		builderSubmission.BuilderSignature,
		builderSubmission.Value,
	)
	BuilderSubmissionHash := sha256.Sum256([]byte(BuilderBid))

	return fmt.Sprintf("%#x", BuilderSubmissionHash)
}
