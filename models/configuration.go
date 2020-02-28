package models

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"reflect"
	"strings"
)

//Client Client
type Client struct {
	Token       string
	RequestURL  string
	RequestFile string
}

type Aria2C struct {
	DownloadDir             string
	SourcesDir              string
	LogDir                  string
	Secret                  string
	Port                    int
	MaxConnectionsPerServer int
	MaxConcurrentDownloads  int
	LogLevel                string
}

//Conf Conf
type Conf struct {
	Title  string
	Client Client
	Aria2C Aria2C
}

//ConfigurationFile файл конфигурации
const ConfigurationFile string = "config.toml"

//ReadConfig функция
func ReadConfig() Conf {
	log.Println("Reading file configuration", ConfigurationFile)
	_, err := os.Stat(ConfigurationFile)
	if err != nil {
		log.Fatal("Config file is missing: ", ConfigurationFile)
	}
	var config Conf
	if _, err := toml.DecodeFile(ConfigurationFile, &config); err != nil {
		log.Fatal("ERROR", err)
	}

	var url = strings.Replace(config.Client.RequestURL, "${token}", config.Client.Token, -1)
	config.Client.RequestURL = url
	var fileRequest = strings.Replace(config.Client.RequestFile, "${token}", config.Client.Token, -1)
	config.Client.RequestFile = fileRequest

	//TODO: MB LATER?
	//var fv = make(map[*reflect.StructField]*reflect.Value)
	//createFieldMap(&config, fv)
	//resolveProperties(&config, fv)
	return config
}

func resolveProperties(c *Conf, fv map[*reflect.StructField]*reflect.Value) {
	for _, v := range fv {
		var neededToResolve = checkString(v.String())
		for _, s := range neededToResolve {
			var val = getVal(s, fv)
			var newVal = strings.Replace(v.String(), "${"+s+"}", val, -1)
			log.Println(newVal)
			log.Println(v.CanAddr())
			log.Println(v.CanSet())

		}
	}
}

func getVal(name string, fv map[*reflect.StructField]*reflect.Value) string {
	for k, v := range fv {
		if strings.EqualFold(k.Name, name) {
			return v.String()
		}
	}
	return ""
}

func checkString(str string) []string {
	var dollar = false
	var leftBrace = false
	var result = make([]string, 0)
	var sb = new(strings.Builder)
	for i, r := range str {
		if r == '$' {
			dollar = true
			continue
		}
		if r == '{' && dollar {
			leftBrace = true
			continue
		}
		if r == ' ' && leftBrace && dollar {
			log.Fatal("Error at %n Right brace expected,", i)
		}
		if r == '}' && dollar && leftBrace {
			if sb.Len() < 1 {
				log.Fatal("Inner text cannot be null")
			}
			result = append(result, sb.String())
			sb.Reset()
			dollar = false
			leftBrace = false
		}
		if dollar && leftBrace {
			sb.WriteRune(r)
		}
	}
	return result
}

func createFieldMap(i interface{}, m map[*reflect.StructField]*reflect.Value) {
	var v = reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var fCount = v.Type().NumField()
	for i := 0; i < fCount; i++ {
		var valueField = v.Field(i)
		var typeField = v.Type().Field(i)
		if typeField.Type.Kind() == reflect.Struct {
			createFieldMap(valueField.Interface(), m)
		} else {
			log.Println(v.String())
			var neededToResolve = checkString(v.String())
			for _, s := range neededToResolve {
				var val = getVal(s, m)
				var newVal = strings.Replace(v.String(), "${"+s+"}", val, -1)
				v.SetString(newVal)
				log.Println(newVal)
				log.Println(v.CanAddr())
				log.Println(v.CanSet())
			}
			m[&typeField] = &valueField
		}
	}
}