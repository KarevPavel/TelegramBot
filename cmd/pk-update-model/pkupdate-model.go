package main

import (
	"golang.org/x/net/html"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"telegram-bot-long-polling/constants"
	"telegram-bot-long-polling/util"
	"text/template"
)

/**
Method to update telegram models
*/

var TypeConverterMap = map[string]string{
	"Integer": "int",
	"String": "string",
	"Boolean": "bool",
	"True": "bool",
	"Integer or String": "string",
	"InputFile or String": "string",
	"Float number": "float64",
	"Float": "float64",
}

const TelegramModelsUrl = "https://core.telegram.org/bots/api"

type StructureDescriptor struct {
	Name, Description string
	Field []StructureField
}

type StructureField struct {
	Name, Description, TypeName string
	IsOptional bool
}

type nodeFunc func(node *html.Node)
type nodePredicate func(node *html.Node) bool

func main() {
	var response, _ = http.Get(TelegramModelsUrl)
	var parsedHtml, _ = html.Parse(response.Body)
	var nodes = make([]*html.Node, 0)
	var nodeAppender nodeFunc = func(node *html.Node) {
		nodes = append(nodes, node)
	}
	doEachNodeIfTrue(parsedHtml, nodeIsH4, nodeAppender)
	var structureDescs = make([]StructureDescriptor, len(nodes))
	var structCreator = func(index int, node *html.Node) {
		var thead = searchEachNode(node, nodeIsTHead)
		var columnNames = getThsDataFromTHead(thead.FirstChild)
		var tbody = searchEachNode(node, nodeIsTBody)
		var rows = getTdsDataFromTBody(tbody.FirstChild)
		var structureDesc = StructureDescriptor{
			Name:        util.ReplaceTabsAndNewLines(getStructName(node)),
			Description: util.ReplaceTabsAndNewLines(getStructDescription(node)),
			Field:       parseRows(rows, columnNames),
		}
		structureDescs[index] = structureDesc
	}

	for index, node := range nodes {
		if isStructDescription(node) {
			structCreator(index, node)
		}
	}

	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}

	var tmpl, err = template.New("template.tmpl").ParseFiles(path.Dir(filePath) + "/template.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Create(path.Dir(filePath) + "/model.go")
	if err := tmpl.Execute(f, structureDescs); err != nil {
		log.Fatal(err)
	}
}

func parseRows(rows []*html.Node, names []string) []StructureField {
	var arrLength = int(math.Ceil(float64(len(rows)) / float64(len(names))))
	var structArr = make([]StructureField, arrLength)
	var structF = StructureField{}
	for i := 0; i < len(rows); i++ {
		var mod = math.Mod(float64(i), float64(len(names)))
		if mod == 0 && i != 0 {
			structArr[int(math.Ceil(float64(i)/float64(len(names)))) - 1] = structF
			structF = StructureField{}
		}
		var colName = names[int(mod)]
		var colValue = getAllTextFromInnerNodes(rows[i].FirstChild)
		switch colName {
		case "Field":
			structF.Name = util.Capitalize(colValue)
		case "Parameter":
			structF.Name = util.Capitalize(colValue)
		case "Description":
			structF.Description = util.ReplaceTabsAndNewLines(colValue)
			if strings.Contains(structF.Description, "Optional") {
				structF.IsOptional = true
			}
		case "Optional":
			if colValue == "Optional" {
				structF.IsOptional = true
			} else {
				structF.IsOptional = false
			}
		case "Type":
			colValue = util.ReplaceTabsAndNewLines(colValue)
			var goType = TypeConverterMap[colValue]
			if goType != "" {
				structF.TypeName = goType
			} else {
				if strings.Contains(colValue, "Array of") {
					structF.TypeName = "[]" + strings.Replace(colValue, "Array of", "", -1)

					if strings.Contains(structF.TypeName, "InputMediaPhoto"){
						structF.TypeName = "[]InputMediaPhoto"
					}

					if strings.Contains(structF.TypeName,"String") {
						structF.TypeName = "[]string"
					}

				} else {
					structF.TypeName = "*" + colValue
				}
			}
		}
	}
	return structArr
}

func nodeIsH4(node *html.Node) bool {
	return isNodeTagEqual(node, "h4")
}

func nodeIsTd(node *html.Node) bool {
	return isNodeTagEqual(node, "td")
}

func nodeIsTh(node *html.Node) bool {
	return isNodeTagEqual(node, "th")
}

func nodeIsTHead(node *html.Node) bool {
	return isNodeTagEqual(node, "thead")
}

func nodeIsTBody(node *html.Node) bool {
	return isNodeTagEqual(node, "tbody")
}

func getThsDataFromTHead(tAny *html.Node) []string {
	var strArr = make([]string, 0)
	var nodeAppender nodeFunc = func(node *html.Node) {
		strArr = append(strArr, node.FirstChild.Data)
	}
	doEachNodeIfTrue(tAny, nodeIsTh, nodeAppender)
	return strArr
}

func getTdsDataFromTBody(tAny *html.Node) []*html.Node {
	var strArr = make([]*html.Node, 0)
	var nodeAppender nodeFunc = func(node *html.Node) {
		strArr = append(strArr, node)
	}
	doEachNodeIfTrue(tAny, nodeIsTd, nodeAppender)
	return strArr
}

func isStructDescription(node *html.Node) bool {
	if node.FirstChild.NextSibling == nil {
		return false
	}
	if node.FirstChild.NextSibling.NextSibling == nil {
		return false
	}

	if node.FirstChild.NextSibling.NextSibling.NextSibling == nil {
		return false
	}

	if node.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild == nil {
		return false
	}

	var structName = node.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	if strings.Contains(structName, " ") || strings.Contains(structName, "-"){
		return false
	}
	return true
}

func getStructName(node *html.Node) string {
	return node.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
}

func getStructDescription(node *html.Node) string {
	return getAllTextFromInnerNodes(node.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild)
}

func isTextNode(node *html.Node) bool {
	return node.Type == html.TextNode
}

func getAllTextFromInnerNodes(node *html.Node) string {
	var sb = new(strings.Builder)
	doEachNodeIfTrue(node, isTextNode, func(node *html.Node) {
		sb.WriteString(node.Data)
	})
	return sb.String()
}

func getAttributeValue(attrs []html.Attribute, attributeKey string) string {
	for _, attr := range attrs {
		if attr.Key == attributeKey {
			return attr.Val
		}
	}
	return constants.EmptyString
}

func isNodeTagEqual(node *html.Node, nodeTag string) bool {
	return node.Data == nodeTag
}

//TODO: Use normal algorithm + goroutins
func doEachNodeIfTrue(root *html.Node, condition nodePredicate, do nodeFunc) {
	for c := root; c != nil; c = c.FirstChild {
		if c.NextSibling != nil {
			doEachSibling(c.NextSibling, condition, do)
		}
		if condition(c) {
			do(c)
		}
	}
}

func returnNodeIfTrue(root *html.Node, condition nodePredicate, do nodeFunc) {
	for c := root; c != nil; c = c.FirstChild {
		if c.NextSibling != nil {
			doEachSibling(c.NextSibling, condition, do)
		}
		if condition(c) {
			do(c)
		}
	}
}


func doEachSibling(node *html.Node, condition nodePredicate, do nodeFunc) {
	for c := node; c != nil; c = c.NextSibling {
		if c.FirstChild != nil {
			doEachNodeIfTrue(c.FirstChild, condition, do)
		}
		if condition(c) {
			do(c)
		}
	}
}

func searchEachNode(node *html.Node, condition nodePredicate) *html.Node {
	for c := node; c != nil; c = c.NextSibling {
		if c.NextSibling != nil {
			if res := searchEachSibling(c.FirstChild, condition); res != nil {
				return res
			}
		}
		if condition(c) {
			return c
		}
	}
	return nil
}

func searchEachSibling(node *html.Node, condition nodePredicate) *html.Node {
	for c := node; c != nil; c = c.NextSibling {
		if c.FirstChild != nil {
			if res := searchEachNode(c.FirstChild, condition); res != nil {
				return res
			}
		}
		if condition(c) {
			return c
		}
	}
	return nil
}