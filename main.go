package main

import (
	"io/ioutil"
	"path/filepath"
	"fmt"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


var (
	config *Config = new(Config)
	app    *fiber.App = fiber.New()

	configFile *string = flag.String("config", "wkr.config", "Config file for Wkr")
	generateHash *bool = flag.Bool("generate-hash", false, "Generate a hash from a given password")
	validateHash *bool = flag.Bool("validate-hash", false, "Validate a password with a hash")
)

func init() {
	flag.Parse()
}

func main() {

	// Generating a hash
	if *generateHash == true {
		generateHashCommand()
		os.Exit(0)
	}

	// Validating a hash
	if *validateHash == true {
		validateHashCommand()
		os.Exit(0)
	}

	// Default program execution

	// Loading config from file
	*configFile, _ = filepath.Abs(*configFile)
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Parsing config
	err = config.Parse(string(content))
	if err != nil {
		log.Fatal(err)
	}

	// Validating config
	err = config.Validate()
	if err != nil {
		log.Fatal(err)
	}


	// Do stuff with config state
	registerAppHandlers()
	app.Listen(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
	
}

func generatePasswordHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func validatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// generateHashCommand is invoked when `wkr -generate-hash` is called
func generateHashCommand() {
	if len(os.Args) != 3 {
		log.Fatal("Argument count wrong!")
	}
	hashed, err := generatePasswordHash(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hashed)
}

// validateHashCommand is invoked when `wkr -validate-hash` is called
func validateHashCommand() {
	if len(os.Args) != 4 {
		log.Fatal("Argument count wrong!")
	}
	password := os.Args[2]
	hash     := os.Args[3]
	if ok := validatePasswordHash(password, hash); !ok {
		fmt.Println("Password and Hash don't match!")
		os.Exit(1)
	}
	fmt.Println("Password and Hash are matching.")
}
