package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/belmadge/otel-zipkin/service-a/tracing"

	"github.com/gin-gonic/gin"
)

type InputRequest struct {
	CEP string `json:"cep"`
}

func isValidCEP(cep string) bool {
	return len(cep) == 8
}

func ForwardToServiceB(cep string) (map[string]interface{}, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", "http://service-b:8081/weather/"+cep, nil)

	tracing.AddTrace(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Service B returned an error")
	}

	var result map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result, nil
}

func HandleInput(c *gin.Context) {
	var input InputRequest
	if err := c.ShouldBindJSON(&input); err != nil || !isValidCEP(input.CEP) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	result, err := ForwardToServiceB(input.CEP)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Service B error"})
		return
	}

	c.JSON(http.StatusOK, result)
}
