package services

import (
	"AiLab1/internal/models"
)

func GetTrueRules(object models.Object, rules []models.Rule) []models.Object {
	// Массив для хранения всех объектов, участвующих в логической цепочке
	objects := make([]models.Object, 0, len(rules))

	// Найдем первое правило, соответствующее объекту
	indexStart := findIndexInRules(object, rules)
	if indexStart >= 0 {
		objects = downToUp(object, rules, indexStart)
	} else {
		objects = appendUniqueObject(objects, object)
	}

	// Проход по правилам для проверки дальнейших действий
	for _, rule := range rules {
		if arrayInArray(rule.Cond, objects) {
			objects = appendUniqueObject(objects, rule.Action)
		}
	}

	// Вызов функции для обратного поиска условий, если есть конечное действие
	objects = reverseChain(objects, rules)

	return objects
}

// reverseChain выполняет обратный анализ цепочки условий, начиная с последнего действия
func reverseChain(objects []models.Object, rules []models.Rule) []models.Object {
	fullChain := objects // Список всех объектов, начнем с переданных объектов

	// Пройдем по каждому объекту и найдем все правила, которые к нему привели
	for _, object := range objects {
		fullChain = recursiveBacktrack(object, fullChain, rules)
	}

	return fullChain
}

// recursiveBacktrack ищет все возможные предыдущие условия для заданного объекта
func recursiveBacktrack(currentObject models.Object, fullChain []models.Object, rules []models.Rule) []models.Object {
	// Пройдем по всем правилам
	for _, rule := range rules {
		// Если текущее действие совпадает с искомым объектом
		if rule.Action.Name == currentObject.Name && rule.Action.Value == currentObject.Value {
			// Добавляем все условия этого правила
			for _, condition := range rule.Cond {
				// Если условие еще не в цепочке, добавляем его и продолжаем рекурсию
				if !inArray(condition, fullChain) {
					fullChain = append(fullChain, condition)
					// Рекурсивно ищем дальше условия для каждого найденного объекта
					fullChain = recursiveBacktrack(condition, fullChain, rules)
				}
			}
		}
	}

	return fullChain
}

func arrayInArray(items []models.Object, array []models.Object) bool {
	count := 0
	for _, item := range items {
		for _, el := range array {
			if el.Name == item.Name && el.Value == item.Value {
				count++
				break
			}
		}
	}
	return count == len(items)
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
		if el.Name == object.Name && el.Value == object.Value {
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
