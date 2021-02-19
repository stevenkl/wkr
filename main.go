package main

import (
	"io/ioutil"
	"path/filepath"
	"fmt"
	"flag"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)


var (
	config *Config = new(Config)

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
	*configFile, _ = filepath.Abs(*configFile)
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = config.Parse(string(content))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}

// generateHashCommand is invoked when `wkr -generate-hash` is called
func generateHashCommand() {
	password := os.Args[2]
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(hashed))
}

// validateHashCommand is invoked when `wkr -validate-hash` is called
func validateHashCommand() {
	password := os.Args[2]
	hash     := os.Args[3]
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("Password and Hash don't match!")
		os.Exit(1)
	}
	fmt.Println("Password and Hash are matching.")
}