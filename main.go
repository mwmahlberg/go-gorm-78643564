package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/avast/retry-go/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()

	log.Println("starting up")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbuser, dbpass, dbhost, dbport, dbname, dbparams)

	db, dbConnError := retry.DoWithData(
		// This function will be called until it returns nil or the number of attempts is reached
		// Note that it returns a *gorm.DB and an error, so we basically can use it easily
		// as a drop in replacement for gorm.Open.
		func() (*gorm.DB, error) {
			return gorm.Open(mysql.Open(dsn), &gorm.Config{})
		},
		// Try 5 times with a 5 second delay between each attempt and only return the last error.
		// See https://pkg.go.dev/github.com/avast/retry-go/v4#Options for more options.
		retry.Attempts(5),
		retry.LastErrorOnly(true),
		retry.Delay(5*time.Second),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("connection attempt %d failed: %s", n, err)
		}))

	if dbConnError != nil {
		panic(fmt.Errorf("failed to connect to database: %w", dbConnError))
	}

	log.Println("Migrating user table")
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(fmt.Errorf("failed to migrate: %w", err))
	}

	log.Println("waiting for signal")
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	<-ctx.Done()
	log.Println("shutting down")
}

// User is a dummy model for testing
type User struct {
	gorm.Model
	Name string
}

var (
	dbuser   string
	dbpass   string
	dbhost   string
	dbport   string
	dbname   string
	dbparams string
)

func init() {
	flag.StringVar(&dbuser, "dbuser", "user", "database user")
	flag.StringVar(&dbpass, "dbpass", "user_password", "database password")
	flag.StringVar(&dbhost, "dbhost", "db", "database host")
	flag.StringVar(&dbport, "dbport", "3306", "database port")
	flag.StringVar(&dbname, "dbname", "dbname", "database name")
	flag.StringVar(&dbparams, "dbparams", "charset=utf8mb4&parseTime=True&loc=Local", "additional database params")
}
