package services

import (
	"AiLab1/internal/models"
)

func GetTrueRules(object models.Object, rules []models.Rule) []models.Object {
	objects := make([]models.Object, 0, len(rules))
	indexStart := findIndexInRules(object, rules)
	if indexStart > 0 {
		objects = downToUp(object, rules, indexStart)
	} else {
		objects = appendUniqueObject(objects, object)
	}
	for _, rule := range rules {
		if arrayInArray(rule.Cond, objects) {
			objects = appendUniqueObject(objects, rule.Action)
		}
	}
	return objects
}

func arrayInArray(items []models.Object, array []models.Object) bool {
	count := 0
	for _, item := range items {
		for _, el := range array {
			if el == item {
				count++
				break
			}
		}
	}
	if count == len(items) {
		return true
	}
	return false
}

func downToUp(object models.Object, rules []models.Rule, index int) []models.Object {
	objects := make([]models.Object, 0, len(rules))
	objects = append(objects, object)
	rulesSlice := rules[:index]
	reversRule := reverseRulesArray(rulesSlice)
	for _, rule := range reversRule {
		if inArray(rule.Action, objects) {
			for _, el := range rule.Cond {
				objects = appendUniqueObject(objects, el)
			}
		}
	}
	return reverseObjectsArray(objects)
}

func reverseRulesArray(array []models.Rule) []models.Rule {
	reverseAr := make([]models.Rule, 0, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		reverseAr = append(reverseAr, array[i])
	}
	return reverseAr
}

func reverseObjectsArray(array []models.Object) []models.Object {
	reverseAr := make([]models.Object, 0, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		reverseAr = append(reverseAr, array[i])
	}
	return reverseAr
}

func findIndexInRules(object models.Object, rules []models.Rule) int {
	for i, rule := range rules {
		if inArray(object, rule.Cond) {
			return i
		}
	}
	return -1
}

func inArray(object models.Object, objects []models.Object) bool {
	for _, el := range objects {
		if el == object {
			return true
		}
	}
	return false
}

func appendUniqueObject(objects []models.Object, object models.Object) []models.Object {
	if inArray(object, objects) {
		return objects
	}
	return append(objects, object)
}
