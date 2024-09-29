package models

import (
	"strings"
)

type Rule struct {
	Text   string
	Cond   []Object
	Action Object
}

func CreateRuleByRow(text string) Rule {
	splitText := strings.Split(text, " то ")
	data := map[string]string{
		"condition": splitText[0],
		"action":    splitText[1],
	}
	data["condition"], _ = strings.CutPrefix(data["condition"], "если")
	for key, value := range data {
		data[key] = strings.TrimSpace(value)
	}
	return Rule{
		Text:   text,
		Cond:   CreateConditionByString(data["condition"]),
		Action: CreateObjectByString(data["action"]),
	}
}

func CreateRulesByText(text string) []Rule {
	rulesRow := strings.Split(text, "\n")
	rules := make([]Rule, 0, len(rulesRow))
	for _, row := range rulesRow {
		rules = append(rules, CreateRuleByRow(row))
	}
	return rules
}

func CreateConditionByString(text string) []Object {
	objectsString := strings.Split(text, " и ")
	for i, _ := range objectsString {
		objectsString[i] = strings.TrimSpace(objectsString[i])
	}
	objects := make([]Object, len(objectsString))
	for i, objectString := range objectsString {
		objects[i] = CreateObjectByString(objectString)
	}
	return objects
}
