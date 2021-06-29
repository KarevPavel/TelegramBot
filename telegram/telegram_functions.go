package telegram

import (
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"bitbucket.org/y4cxp543/telegram-bot/interfaces"
	"bitbucket.org/y4cxp543/telegram-bot/telegram/models"
	"bitbucket.org/y4cxp543/telegram-bot/util"
	"net/http"
	"strconv"
	"strings"
)

type TFunctions struct {
	Url         string
	FileRequest string
}

func NewTFunctions(Url string, FileRequest string) interfaces.ITelegramFunctions {
	return &TFunctions{
		FileRequest: FileRequest,
		Url:         Url,
	}
}

func (tFunc *TFunctions) GetMe() (models.User, error) {
	url := util.ReplaceMethod(tFunc.Url, constants.Method, constants.GetMe)
	var answer = models.User{}
	_ = util.DoGet(url, &answer)
	return answer, nil
}

func (tFunc *TFunctions) GetFile(fileId string) (models.File, error) {
	url := util.ReplaceMethod(tFunc.Url, constants.Method, constants.GetFile)
	var answer = new(models.File)
	_ = util.DoPost(url, models.GetFile{
		FileId: fileId,
	}, answer)
	return *answer, nil
}

func (tFunc *TFunctions) DownloadFile(filePath string) []byte {
	var url = util.Replace(tFunc.FileRequest, "filePath", filePath)
	var response, _ = http.Get(url)
	var byteArr = util.GetBytes(response)
	return byteArr
}


func (tFunc *TFunctions) GetUpdates(query models.GetUpdates, response chan []models.Update){
	url := util.ReplaceMethod(tFunc.Url, constants.Method, constants.GetUpdates)
	var builder strings.Builder
	builder.WriteString(url)
	var isFirst = false
	if query.Offset != 0 {
		util.AddQueryParam(&builder, &isFirst, constants.Offset, strconv.Itoa(query.Offset))
	}
	if query.Timeout != 0 {
		util.AddQueryParam(&builder, &isFirst, constants.Timeout, strconv.Itoa(query.Timeout))
	}
	if query.Limit != 0 {
		util.AddQueryParam(&builder, &isFirst, constants.Limit, strconv.Itoa(query.Limit))
	}
	if len(query.AllowedUpdates) > 0 {
		for _, allowedUpdate := range query.AllowedUpdates {
			util.AddQueryParam(&builder, &isFirst, constants.AllowedUpdates, allowedUpdate)
		}
	}
	var answer []models.Update
	_ = util.DoGet(builder.String(), &answer)
	response <- answer
}

func (tFunc *TFunctions) SendMessage(request models.SendMessage) (models.Message, error) {
	url := util.ReplaceMethod(tFunc.Url, constants.Method, constants.SendMessage)
	message := models.Message{}
	if err := util.DoPost(url, request, &message); err != nil {
		return message, err
	}
	return message, nil
}

func (tFunc *TFunctions) SendPoll(poll models.SendPoll) (models.Message, error) {
	url := util.ReplaceMethod(tFunc.Url, constants.Method, constants.SendPoll)
	message := models.Message{}
	if err := util.DoPost(url, poll, &message); err != nil {
		return message, err
	}
	return message, nil
}
