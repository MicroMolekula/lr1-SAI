package services

import (
	"AiLab1/internal/models"
	"AiLab1/internal/responses"
)

func GetMapObject(text string) []responses.ObjectValuesResponse {
	rules := models.CreateRulesByText(text)
	objects := getAllObjectsByRules(rules)
	names := getNamesArray(objects)
	result := make([]responses.ObjectValuesResponse, 0, 10)
	for _, name := range names {
		tmp := responses.ObjectValuesResponse{
			Name:   name,
			Values: getValuesForObject(name, objects),
		}
		result = append(result, tmp)
	}
	return result
}

func getAllObjectsByRules(rules []models.Rule) []models.Object {
	objects := make([]models.Object, 0, len(rules))
	for _, rule := range rules {
		for _, el := range rule.Cond {
			if !inArray(el, objects) {
				objects = append(objects, el)
			}
		}
		if !inArray(rule.Action, objects) {
			objects = append(objects, rule.Action)
		}
	}
	return objects
}

func getNamesArray(objects []models.Object) []string {
	result := make([]string, 0, len(objects))
	for _, object := range objects {
		if !inNameObjectArray(object.Name, result) {
			result = append(result, object.Name)
		}
	}
	return result
}

func inNameObjectArray(name string, stringArray []string) bool {
	for _, el := range stringArray {
		if el == name {
			return true
		}
	}
	return false
}

func getValuesForObject(name string, objects []models.Object) []string {
	values := make([]string, 0, len(objects))
	for _, object := range objects {
		if object.Name == name {
			values = append(values, object.Value)
		}
	}
	return values
}
