package utils

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"strings"
)

func SaveBytesToFile(filePath, fileName string, byteArr []byte) string {
	var builder = new(strings.Builder)
	_ = os.MkdirAll(filePath, os.ModeDir)
	builder.WriteString(filePath)
	if !EndsWithPathSeparator(filePath){
		builder.WriteString(string(os.PathSeparator))
	}
	builder.WriteString(fileName)
	filePath = builder.String()
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	_, _ = file.Write(byteArr)
	return filePath
}

func EndsWithPathSeparator(filePath string) bool {
	if len(filePath) < 1 {
		return true
	}

	if filePath[len(filePath) - 1] == os.PathSeparator {
		return true
	}
	return false
}

//ReadConfig функция
func ReadConfig(configFile string, typ interface{}) interface{} {
	log.Println("Reading file configuration", configFile)
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing: ", configFile)
	}
	if _, err := toml.DecodeFile(configFile, typ); err != nil {
		log.Fatal("ERROR", err)
	}
	return typ
}