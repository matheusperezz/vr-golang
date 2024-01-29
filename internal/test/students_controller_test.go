package test

import (
	"bytes"
	"encoding/json"
	"io"
	"main/internal/models"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
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

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateStudent(t *testing.T) {
	student := models.Student{
		Name: "Test Student",
	}

	jsonStudent, _ := json.Marshal(student)
	resp, err := http.Post("http://localhost:8080/students", "application/json", bytes.NewBuffer(jsonStudent))
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}

	var createdStudent models.Student
	json.NewDecoder(resp.Body).Decode(&createdStudent)

	createdStudent.Name = "Updated Test Student"
	endpoint := "http://localhost:8080/students/" + strconv.Itoa(int(createdStudent.ID))
	jsonStudent, _ = json.Marshal(createdStudent)
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(jsonStudent))
	if err != nil {
		t.Fatalf("An error occured %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err = client.Do(req)
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
