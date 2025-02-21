package main

import (
	// "fmt"
	"strings"
)

type Target interface {
	getCsvData() string
}

type NewLibrary struct{}

func (l NewLibrary) getJsonData() []map[string]string {
	return []map[string]string{
		{"data1": "json_dataA", "data2": "json_dataB"},
		{"data1": "json_dataC", "data2": "json_dataD"},
	}
}

type JsonToCsvAdapter struct {
	adaptee *NewLibrary
}

func (a *JsonToCsvAdapter) getCsvData() string {
	jsonData := a.adaptee.getJsonData()
	if len(jsonData) == 0 {
		return ""
	}

	var headers []string
	for key := range jsonData[0] {
		headers = append(headers, key)
	}
	headerLine := strings.Join(headers, ",") + "\n"

	var bodyLines []string
	for _, row := range jsonData {
		var values []string
		for _, key := range headers {
			values = append(values, row[key])
		}
		line := strings.Join(values, ",")
		bodyLines = append(bodyLines, line)
	}
	body := strings.Join(bodyLines, "\n")

	return headerLine + body
}

// func main() {
// 	adaptee := NewLibrary{}
// 	fmt.Println("=== Adapteeが提供するデータ ===")
// 	fmt.Println(adaptee.getJsonData())

// 	fmt.Println("")
// 	adapter := JsonToCsvAdapter{&adaptee}
// 	fmt.Println("=== Adapterに変換されたデータ ===")
// 	fmt.Println(adapter.getCsvData())
// }
