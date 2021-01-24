package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go/build"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

// function for loading configuration file
func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(AppFilePath(filename))
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	if err != nil {
		fmt.Println("decoding error")
	}

	return config, err
}

func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// rendering templates
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func AppFilePath(fpath string) string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return filepath.Join(gopath, "src/ebus/"+fpath)
}

func DBCon(config Config) (*gorm.DB, error) {
	// fmt.Println(config)
	deploy := config.Deployment
	if deploy == "live" {
		db_config := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s 
			sslmode=disable`,
			config.LiveDB.Host,
			config.LiveDB.User,
			config.LiveDB.Password,
			config.LiveDB.DbName)
		db, err := gorm.Open("postgres", db_config)
		return db, err
	} else if deploy == "test" {
		db_config := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s 
			sslmode=disable`,
			config.TestDB.Host,
			config.TestDB.User,
			config.TestDB.Password,
			config.TestDB.DbName)
		db, err := gorm.Open("postgres", db_config)
		return db, err
	}
	return nil, &ConfigError{}
}

// random generators

var DefaultGenerator = func() string {
	id := uuid.NewV4()
	return id.String()
}

// Generate a random string of A-Z chars with len = l
func RandomString(l int) string {
	var pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
