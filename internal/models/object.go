package models

import (
	"strings"
)

type Object struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func CreateObjectByString(text string) Object {
	data := strings.Split(text, "=")
	for i := range data {
		data[i] = strings.TrimSpace(data[i])
	}
	return Object{
		Name:  data[0],
		Value: data[1],
	}
}
