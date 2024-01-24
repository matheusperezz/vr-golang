package test

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"io"
	"main/api/models"
	"net/http"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	student := models.Student{
		Name: "Test Student",
	}

	jsonStudent, _ := json.Marshal(student)

	resp, err := http.Post("http://localhost:8080/students", "application/json", bytes.NewBuffer(jsonStudent))
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("An error occured %v", err)
		}
	}(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateStudent(t *testing.T) {
	student := models.Student{
		Name: "Updated Student",
	}

	jsonStudent, _ := json.Marshal(student)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/students/1", bytes.NewBuffer(jsonStudent))
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
