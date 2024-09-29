package main

import (
	"AiLab1/internal/models"
	"AiLab1/internal/parser"
	"AiLab1/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	fileName = "base_knowledge_5.txt"
)

func handlerObjectsValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	data, err := parser.GetFileData(fileName)
	if err != nil {
		panic(err)
	}
	res := services.GetMapObject(data)
	jsonData, err := json.Marshal(res)
	if err != nil {
		return
	}
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

func handlerResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if r.Method == "POST" {
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		requestData := models.Object{}
		json.Unmarshal(requestBody, &requestData)
		data, err := parser.GetFileData(fileName)
		rules := models.CreateRulesByText(data)
		if err != nil {
			panic(err)
		}
		responseData := services.GetTrueRules(requestData, rules)
		jsonData, err := json.Marshal(responseData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(jsonData)
	}
}

func main() {
	http.HandleFunc("/api/objects/values/", handlerObjectsValue)
	http.HandleFunc("/api/objects/", handlerResult)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func cli() {
	data, err := parser.GetFileData("base_knowledge.txt")
	if err != nil {
		panic(err)
	}
	var name string
	var value string
	fmt.Print("Имя: ")
	fmt.Scan(&name)
	fmt.Print("Значение: ")
	fmt.Scan(&value)
	fmt.Println("__________")
	rules := models.CreateRulesByText(data)
	objectOne := models.Object{
		Name:  name,
		Value: value,
	}
	for _, el := range services.GetTrueRules(objectOne, rules) {
		fmt.Println(fmt.Sprintf("%s = %s", el.Name, el.Value))
	}
}
