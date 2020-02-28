package torrentz2

import (
	"net/http"
	"regexp"
	"strings"
	"telegram-bot-long-polling/constants"
	"telegram-bot-long-polling/util"
)


type Result struct {
	Link, Name, Age, Size string
}

func Search(condition string) []Result {
	condition = strings.ReplaceAll(condition, constants.Space, "+")
	response,_ := http.Get(constants.SearchOrderByPeers + condition)
	result:= string(util.GetBytes(response)[:])
	regex := regexp.MustCompile("href=(?P<Link>\\/[a-zA-Z0-9]+)>(?P<Name>[а-яА-Яa-zA-Z0-9\\s-\\(\\)\\.\\_\\[\\]]+)</a>.*title=\\d+>(?P<Age>[0-9]+[\\s|year|month|day|D]*s?)</span><span>(?P<Size>\\d+[\\sGB|MB|KB]+)")
	submatches := regex.FindAllStringSubmatch(result , -1)
	var answer = make([]Result, len(submatches))
	for index, submatch := range submatches {
		var result = Result{}
		for i, name := range regex.SubexpNames() {
			switch name {
			case "Link":
				result.Link = submatch[i]
			case "Name":
				result.Name = submatch[i]
			case "Age":
				result.Age = submatch[i]
			case "Size":
				result.Size = submatch[i]
			default:
				continue
			}
		}
		answer[index] = result
	}
	return answer
}