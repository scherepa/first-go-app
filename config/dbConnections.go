package config

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// lets add .env
	"os"
	// This is the driver we just "got"
	//In Go, _ means “import for side effects only”
	_ "github.com/go-sql-driver/mysql"
)

type DBConnConfig struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
	driver   string
}

// Global registry for database connections
var DBRegistry = make(map[string]*sql.DB)

// InitConnection sets up a connection and stores it in the map
func InitConnection(name string, driver string) {
	dsn, err := LoadDsn(name, driver)
	if err != nil {
		log.Fatalf("Could not load DSN for %s: %v", name, err)
	}
	// should be cases for different types
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatalf("Error opening %s: %v", name, err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Connection %s failed: %v", name, err)
	}
	// Store it in our "Connections Registry"
	DBRegistry[name] = db
}

func PingConn(name string) {
	err := DBRegistry[name].Ping()
	// The "Ping" checks if the bridge to db is actually working
	if err != nil {
		log.Fatalf("Connection '%s' Failed! Is Docker running? Error: %v", name, err)
	}
	fmt.Printf("🚀 Ping SUCCESS! Go engine is talking to '%s'\n", name)
}

func LoadDsn(key string, driver string) (string, error) {
	keyUp := strings.ToUpper(key)
	// Read env vars
	conf := DBConnConfig{
		user:     os.Getenv(keyUp + "_USER"),
		password: os.Getenv(keyUp + "_ROOT_PASSWORD"),
		host:     os.Getenv(keyUp + "_HOST"),
		port:     os.Getenv(keyUp + "_PORT"),
		dbname:   os.Getenv(keyUp + "_DATABASE"),
		driver:   driver,
	}
	// Format: user:password@tcp(127.0.0.1:3307)/dbname
	// dsn := "root:your_secure_password@tcp(127.0.0.1:3307)/my_app"
	// DSN format: user:password@tcp(host:port)/dbname
	// later on mabe case postgre
	// case
	switch conf.driver {
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
			conf.user,
			conf.password,
			conf.host,
			conf.port,
			conf.dbname,
		), nil
	default:
		return "", fmt.Errorf("unsupported db driver")
	}
}
