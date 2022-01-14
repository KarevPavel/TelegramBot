package util

import (
	"bitbucket.org/y4cxp543/aria2c"
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/telegram/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func GetBytes(response *http.Response) []byte {
	byteArr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
		return nil
	}
	return byteArr
}

func GetGid(resp *aria2c.Response) string {
	var gid = ""
	if resp.Params != nil {
		log.Println(reflect.TypeOf(resp.Params).Kind())
		if reflect.TypeOf(resp.Params).Kind() == reflect.Slice {
			var arr = resp.Params.([]interface{})
			if len(arr) == 1 && arr[0] != nil && reflect.TypeOf(arr[0]).Kind() == reflect.Map {
				var kv = arr[0].(map[string]interface{})
				var val = kv["gid"]
				if reflect.TypeOf(val).Kind() == reflect.String {
					gid = val.(string)
				}
			}
		}
	}
	return gid
}

func ReplaceMethod(source, condition string, method constants.TelegramMethods) string {
	return Replace(source, condition, method.String())
}

func Replace(source, condition, method string) string {
	return strings.ReplaceAll(source, `${` + condition + `}`, method)
}

func AddQueryParam(builder *strings.Builder, isFirst *bool, queryParam constants.QueryParams, paramValue string){
	if *isFirst {
		builder.WriteString(fmt.Sprintf("&%s=%s", queryParam, paramValue))
	} else {
		*isFirst = true
		builder.WriteString(fmt.Sprintf("?%s=%s", queryParam, paramValue))
	}
}


func Guid() string {
	var uuid, _ = uuid.NewV4()
	return uuid.String()
}

func AddSpacesBetweenStrings(stringArr ...string) string {
	var builder = new(strings.Builder)
	for index, str := range stringArr {
		if index != 0 {
			builder.WriteString(constants.Space)
		}
		builder.WriteString(str)
	}
	return builder.String()
}

func DoPost(url string, request interface{}, object interface{}) error {
	var marshal, err = json.Marshal(request)
	if err != nil {
		log.Panic(err)
		return err
	}
	log.Println("Request: " + string(marshal))
	response, err := http.Post(url, constants.JSONContentType, bytes.NewReader(marshal))
	apiResponse := models.APIResponse{}
	_ = UnmarshalToType(response, &apiResponse)
	result, _ := apiResponse.Result.MarshalJSON()
	if err != nil {
		log.Panic(err)
		return err
	}
	//return UnmarshalToType(response, object)
	return json.Unmarshal(result, &object)
}

func DoGet(url string, object interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
		return err
	}
	apiResponse := models.APIResponse{}
	_ = UnmarshalToType(response, &apiResponse)
	result, _ := apiResponse.Result.MarshalJSON()
	return json.Unmarshal(result, &object)
}


func UnmarshalToType(response *http.Response, object interface{}) error {
	bytesArr := GetBytes(response)
	log.Println("Response: " + string(bytesArr[:]))
	return json.Unmarshal(bytesArr, object)
}

func Capitalize(source string) string {
	if strings.Contains(source, "_"){
		var sb = new(strings.Builder)
		var separated = strings.Split(source, "_")
		for _, separate := range separated {
			sb.WriteString(strings.Title(separate))
		}
		return sb.String()
	} else {
		return strings.Title(source)
	}
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
		return constants.EmptyString
	}
	_, _ = file.Write(byteArr)
	return filePath
}
