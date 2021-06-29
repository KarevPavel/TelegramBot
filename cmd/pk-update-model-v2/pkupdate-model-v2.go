package main

import (
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/util"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"text/template"
)

const TelegramModelsUrl = "https://core.telegram.org/bots/api"

type StructureDescriptor struct {
	Name, Description string
	Field             []StructureField
}

type StructureField struct {
	Name, Description, SerializableName, TypeName string
	IsOptional                                    bool
}

func main() {
	var response, _ = http.Get(TelegramModelsUrl)
	var doc, error = goquery.NewDocumentFromReader(response.Body)
	if error != nil {
		log.Panic(error)
	}

	var structureDescriptions = make([]StructureDescriptor, 0)
	doc.Find("html body div div div div div h4").Each(func(i int, selection *goquery.Selection) {
		var structureName = strings.Title(selection.Text())
		if !strings.Contains(structureName, constants.Space) {
			var h4 = selection.NextUntil("h4")
			var structureDescription = h4.Closest("p").First().Text()
			var colNames = h4.Find("table thead tr th")
			var colValues = h4.Find("table tbody tr")
			if structureDescription != constants.EmptyString &&
				colNames.Size() > 0 &&
				colValues.Size() > 0 {
				var structDesc = StructureDescriptor{
					Name:        structureName,
					Description: structureDescription,
					Field:       createStructFields(colNames, colValues),
				}
				structureDescriptions = append(structureDescriptions, structDesc)
			}
		}
	})

	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}

	var tmpl, err = template.New("template.tmpl").ParseFiles(path.Dir(filePath) + "/template.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Create(path.Dir(filePath) + "/model.go")
	if err := tmpl.Execute(f, structureDescriptions); err != nil {
		log.Fatal(err)
	}
	log.Println("Alright, saved to " + filePath)
}

func createStructFields(names *goquery.Selection, values *goquery.Selection) []StructureField {
	var structureFields = make([]StructureField, values.Size())
	values.Each(func(i int, selection *goquery.Selection) {
		var structureField = StructureField{}
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			names.Text()
			var columnName = names.Get(i).FirstChild.Data
			switch columnName {
			case "Field", "Parameter":
				structureField.SerializableName = selection.Text()
				structureField.Name = util.Capitalize(structureField.SerializableName)
			case "Type":
				var documentationType = selection.Text()
				if goType := commonTypeToGoType(documentationType); goType != reflect.Invalid {
					structureField.TypeName = goType.String()
				} else {
					structureField.TypeName = tryHard(documentationType, &structureFields)
				}
			case "Description":
				structureField.Description = selection.Text()
				if strings.Contains(structureField.Description, "Optional") {
					structureField.IsOptional = true
				}
			case "Required":
				if strings.Contains(selection.Text(), "Optional") {
					structureField.IsOptional = true
				} else {
					structureField.IsOptional = false
				}
			}
		})
		structureFields[i] = structureField
	})
	return structureFields
}

func commonTypeToGoType(commonType string) reflect.Kind {
	switch commonType {
	case "String":
		return reflect.String
	case "Boolean":
		return reflect.Bool
	case "Integer":
		return reflect.Int
	case "Float", "Float number":
		return reflect.Float32
	case "Integer or String":
		return reflect.String
	case "True":
		return reflect.Bool
	}
	return reflect.Invalid
}

//TODO: Too stupid
func tryHard(docType string, fields *[]StructureField) string {

	if strings.Contains(docType, "Array of ") {
		var pureType = strings.Replace(docType, "Array of ", "", -1)
		if goType := commonTypeToGoType(pureType); goType != reflect.Invalid {
			return "[]" + goType.String()
		} else {
			if pureType == "InputMediaPhoto and InputMediaVideo" {
				return "[]InputMediaVideo"
			}

			if pureType == "PassportElementError" {
				return "[]PassportElementErrorDataField"
			}

			if pureType == "InlineQueryResult" {
				return "[]InlineQueryResultCachedAudio"
			}
			return "[]" + pureType
		}
	} else {
		if strings.Contains(docType, "InputFile") {
			return "string"
		}

		if strings.EqualFold(docType, "InputMedia") {
			*fields = append(*fields,
				createOptionalStructureField("InputMediaAnimation",
					"Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.",
					"media",
					"*InputMediaAnimation"))

			*fields = append(*fields,
				createOptionalStructureField("InputMediaAudio",
					"Represents an audio file to be treated as music to be sent.",
					"media",
					"*InputMediaAudio"))

			*fields = append(*fields,
				createOptionalStructureField("InputMediaPhoto",
					"Represents a photo to be sent.",
					"media",
					"*InputMediaPhoto"))

			*fields = append(*fields,
				createOptionalStructureField("InputMediaVideo",
					"Represents a video to be sent.",
					"media",
					"*InputMediaVideo"))

			return "InputMediaDocument"

		}

		if strings.EqualFold(docType, "InlineQueryResult") {

			return "InlineQueryResultVoice"
			/*TODO: Add this QueryRequestTypes:
						InlineQueryResultCachedAudio
						InlineQueryResultCachedDocument
						InlineQueryResultCachedGif
						InlineQueryResultCachedMpeg4Gif
						InlineQueryResultCachedPhoto
						InlineQueryResultCachedSticker
						InlineQueryResultCachedVideo
						InlineQueryResultCachedVoice
						InlineQueryResultArticle
						InlineQueryResultAudio
						InlineQueryResultContact
						InlineQueryResultGame
						InlineQueryResultDocument
						InlineQueryResultGif
						InlineQueryResultLocation
						InlineQueryResultMpeg4Gif
						InlineQueryResultPhoto
						InlineQueryResultVenue
						InlineQueryResultVideo
						InlineQueryResultVoice
			*/
		}

		if strings.EqualFold(docType, "InputMessageContent") {
			*fields = append(*fields,
				createOptionalStructureField("InputTextMessageContent",
					"Represents the content of a text message to be sent as the result of an inline query",
					"input_message_content",
					"*InputTextMessageContent"))

			*fields = append(*fields,
				createOptionalStructureField("InputLocationMessageContent",
					"Represents the content of a text message to be sent as the result of an inline query",
					"input_message_content",
					"*InputLocationMessageContent"))

			*fields = append(*fields,
				createOptionalStructureField("InputVenueMessageContent",
					"Represents the content of a text message to be sent as the result of an inline query",
					"input_message_content",
					"*InputVenueMessageContent"))

			return "*InputContactMessageContent"
		}

		if strings.EqualFold(docType, "CallbackGame") {
			return "string //dont know what is that"
		}

		if strings.Contains(docType, "InlineKeyboardMarkup") {
			//TODO: There is much more markups, soo add them
			return "InlineKeyboardMarkup"
		}

		return "*" + docType
	}
}

func createOptionalStructureField(name, description, serializableName, typeName string) StructureField {
	return StructureField{
		Name:             name,
		Description:      description,
		SerializableName: serializableName,
		TypeName:         typeName,
		IsOptional:       true,
	}
}
