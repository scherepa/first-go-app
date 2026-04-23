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
	sslMode  string // SSL mode for the database connection.
	timeZone string // Time zone for the database.
}

// feel the despereta need of DB Manager
// and separating private and public
// Manager should be available
// but configurations should be available only to DB Manager
// It will have all the methods and map of connections as private property

type Connection interface {
	LoadDsn(name string) (string, error)
	Connect(dsn string) (*sql.DB, error)
}

func (conn MySQLConnection) Connect(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}

func (conn PostgresConnection) Connect(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}

// Global registry for database connections
var DBRegistry = make(map[string]*sql.DB)
var DBConnConfigs = make(map[string]DBConnConfig)
var ConnectionNames = map[string]string{
	"mysql": "mysql",
}

func loadConnectionConfig(name string) {
	driver := ConnectionNames[name]
	keyUp := strings.ToUpper(name)
	mode := os.Getenv(keyUp + "_DB_MODE")
	tz := os.Getenv(keyUp + "_DB_TIMEZONE")
	if tz == "" {
		tz = "Local"
	}
	conf := DBConnConfig{
		user:     os.Getenv(keyUp + "_USER"),
		password: os.Getenv(keyUp + "_ROOT_PASSWORD"),
		host:     os.Getenv(keyUp + "_HOST"),
		port:     os.Getenv(keyUp + "_PORT"),
		dbname:   os.Getenv(keyUp + "_DATABASE"),
		driver:   driver,
		sslMode:  mode,
		timeZone: tz,
	}
	DBConnConfigs[name] = conf
}

// MySQLConnection implements DatabaseConnection for MySQL.
type MySQLConnection struct {
	Config *DBConnConfig
}

// factory GO way
func (conn MySQLConnection) LoadDsn(name string) (string, error) {
	conf := DBConnConfigs[name]
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=%s",
		conf.user,
		conf.password,
		conf.host,
		conf.port,
		conf.dbname,
		conf.timeZone,
	), nil
}

// PostgresConnection implements DatabaseConnection for PostgreSQL.
type PostgresConnection struct {
	Config *DBConnConfig
}

// Connect connects to a PostgreSQL database and returns a GORM DB instance.
func (p PostgresConnection) LoadDsn(name string) (string, error) {
	conf := DBConnConfigs[name]
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.host,
		conf.user,
		conf.password,
		conf.dbname,
		conf.port,
		conf.sslMode,
		conf.timeZone,
	), nil
}

// InitConnection sets up a connection and stores it in the map
func InitConnection(name string, driver string) (*sql.DB, error) {
	loadConnectionConfig(name)
	conf, ok := DBConnConfigs[name]
	if !ok {
		return nil, fmt.Errorf("No config defined for db connecton %s", name)
	}
	var dbConn Connection
	switch conf.driver {
	case "mysql":
		dbConn = &MySQLConnection{Config: &conf}
	case "postgres":
		dbConn = &PostgresConnection{Config: &conf}
	default:
		return nil, fmt.Errorf("Unsupported database type: %#v", conf.driver)
	}
	dsn, err := dbConn.LoadDsn(name)
	if err != nil {
		return nil, fmt.Errorf("Could not load DSN for %s: %v", name, err)
	}
	// should be cases for different types
	db, err := dbConn.Connect(dsn) //sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("Error opening %s: %v", name, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Connection %s failed: %v", name, err)
	}
	// Store it in "Connections Registry"
	DBRegistry[name] = db
	return db, nil
}

func PingConn(name string) {
	err := DBRegistry[name].Ping()
	// The "Ping" checks if the bridge to db is actually working
	if err != nil {
		log.Fatalf("Connection '%s' Failed! Is Docker running? Error: %v", name, err)
	}
	fmt.Printf("🚀 Ping SUCCESS! Go engine is talking to '%s'\n", name)
}

/*func LoadDsn(key string, driver string) (string, error) {
keyUp := strings.ToUpper(key)*/
// Read env vars
/*conf := DBConnConfig{
	user:     os.Getenv(keyUp + "_USER"),
	password: os.Getenv(keyUp + "_ROOT_PASSWORD"),
	host:     os.Getenv(keyUp + "_HOST"),
	port:     os.Getenv(keyUp + "_PORT"),
	dbname:   os.Getenv(keyUp + "_DATABASE"),
	driver:   driver,
}*/
// Format: user:password@tcp(127.0.0.1:3307)/dbname
// dsn := "root:your_secure_password@tcp(127.0.0.1:3307)/my_app"
// DSN format: user:password@tcp(host:port)/dbname
// later on maybe case postgre
// case
//switch conf.driver {
//case "mysql":
//return fmt.Sprintf(
//"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
//conf.user,
//conf.password,
//conf.host,
//conf.port,
//conf.dbname,
//), nil
//default:
//	return "", fmt.Errorf("unsupported db driver")
//}
//}
