package config

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// LoadEnv loads env vars from .env.test
func loadEnv() {
	// Only for local tests - Env vars for actual runtime
	// Source: https://github.com/joho/godotenv/issues/43#issuecomment-503183127
	re := regexp.MustCompile(`^(.*src)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	godotenv.Load(string(rootPath) + `/../.env.test`)
	//if err != nil {
	//	log.WithFields(log.Fields{
	//		"cause": err,
	//		"cwd":   cwd,
	//	}).Fatal("Problem loading .env.test file")
	//
	//	os.Exit(-1)
	//}
}

func ReadTestEnvironment() {
	loadEnv()
	ReadEnvironment()
}
