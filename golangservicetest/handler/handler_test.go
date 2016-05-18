package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golangservicetest/golangservicetest/protobuf"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func Test_handler(t *testing.T) {
	stateHandle := State()
	req, _ := http.NewRequest("GET", "", nil)

	w := httptest.NewRecorder()
	stateHandle.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK, "Home page didn't return %v")

}

func Test_handlerPost(t *testing.T) {
	stateHandle := StatePost()
	var err error = nil
	var jsonStr = []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)
	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	stateHandle.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK, "Home page didn't return %v")

	stateService := new(StateService)
	buf := new(bytes.Buffer)
	buf.ReadFrom(w.Body)
	responseBody := buf.Bytes()
	err = json.Unmarshal(responseBody, stateService)

	if err != nil {
		fmt.Println("Error", "", err)

	} else {
		// It manege to catch the name of the machine(HostName)

		assert.Equal(t, "Tuesday", stateService.Day)
		assert.Equal(t, "INNOVA4J-13", stateService.HostName)
		assert.Equal(t, "6:42 AM", stateService.Hour)
		assert.Equal(t, "ruote", stateService.LogPath)
		assert.Equal(t, "ok", stateService.Status)
	}

}

func Test_handlerPostProto(t *testing.T) {
	stateHandle := StatePost()
	var err error = nil

	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
		Union: &example.Test_Name{"fred"},
	}
	data, err := proto.Marshal(test)

	req, _ := http.NewRequest("POST", "", bytes.NewBuffer(data))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/x-protobuf")

	w := httptest.NewRecorder()
	stateHandle.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK, "Home page didn't return %v")

	person := new(example.Person)
	buf := new(bytes.Buffer)
	buf.ReadFrom(w.Body)
	responseBody := buf.Bytes()
	err = proto.Unmarshal(responseBody, person)

	if err != nil {
		fmt.Println("Error Unmarshal protobuf", "", err)

	} else {
		// It manege to catch the name of the machine(HostName)

		assert.Equal(t, int32(1234), *person.Id)
		assert.Equal(t, "John Doe", *person.Name)
		assert.Equal(t, "jdoe@example.com", *person.Email)

	}

}
