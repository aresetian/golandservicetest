package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golangservicetest/golangservicetest/log"
	"golangservicetest/golangservicetest/protobuf"
	"net/http"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
)

/*
Description :

   The main.go and mappings(handler,gi) separation are posibble due to the Http packages
   are global to all the context of Golang. The goal of this method is define all the mappings
   between URL and handlers.


Parameters :
       pContext :  It has the coxtext of the project. this context is used to concact it to
       the URL

Return :
       No data to return.

*/
func SetUpHandlers(pContext string) {
	//It define the mapping between the URL("/"+pContext+"/State") y the handler(State)
	http.HandleFunc("/"+pContext+"/State", State())
	http.HandleFunc("/"+pContext+"/StatePost", StatePost())
}

/*
   This struct contain the definition of the object that will be used on the respond.
*/
type StateService struct {
	Status   string
	Day      string
	Hour     string
	HostName string
	LogPath  string
}

type RequestPost struct {
	Page   int
	Fruits []string
}

/*

Description :
    Method that create a JSON Object and represents the data of an  StateService Struct.


	Changes: I changed the wait to manage handlers how is said on
	http://www.markjberger.com/testing-web-apps-in-golang/


Parameters :
    pResponse http.ResponseWriter :  It contains the response that it will be send to the final user.
    pRequest *http.Request :         It contains the request od the user.

Return :
       No data to return.

*/

func State() http.HandlerFunc {

	return http.HandlerFunc(func(pResponse http.ResponseWriter, r *http.Request) {
		// It manege to catch the name of the machine(HostName)
		hostname, err := os.Hostname()
		if err != nil {
			// In case does not catch the Host name
			log.Error("It has occur an Error when the HostName", err)
		}
		fmt.Println("Executing State Method")
		state := StateService{"ok", "Tuesday", "6:42 AM", hostname, "ruote"}
		js, _ := json.Marshal(state)
		pResponse.Header().Set("Content-Type", "application/json")
		pResponse.Write(js)
	})
	/*
			// It manege to catch the name of the machine(HostName)
		    hostname, err := os.Hostname()
		    if err != nil {
		      // In case does not catch the Host name
		      log.Error("It has occur an Error when the HostName", err)
		    }
		    state := StateService{"ok","Tuesday","6:41 AM",hostname,"ruote"}
		    js, _ := json.Marshal(state)
		    pResponse.Header().Set("Content-Type", "application/json")
		    pResponse.Write(js)
	*/
}

func StatePost() http.HandlerFunc {

	return http.HandlerFunc(func(pResponse http.ResponseWriter, pRequest *http.Request) {

		requestPost := new(RequestPost)

		contentType := pRequest.Header.Get("Content-Type")
		var err error = nil
		buf := new(bytes.Buffer)
		buf.ReadFrom(pRequest.Body)

		requestBody := buf.Bytes()
		fmt.Println("Executing State Method POST")
		// Si el cuerpo de la peticion contiene informacion se procesa, de lo contrario no se procesaran  peticiones vacias
		if len(requestBody) > 0 {

			if contentType == "application/x-protobuf" || strings.Contains(contentType, "application/x-protobuf") {
				fmt.Println("Executing Error Protobuf ")
				p := &example.Person{
					Id:    proto.Int32(1234),
					Name:  proto.String("John Doe"),
					Email: proto.String("jdoe@example.com"),
				}
				data, err := proto.Marshal(p)
				if err != nil {
					// In case does not catch the Host name
					log.Error("It has occur an Error when the proto.Marshal", err)
				}
				pResponse.Header().Set("Content-Type", "application/x-protobuf")
				pResponse.Write(data)

			} else if contentType == "application/json" || strings.Contains(contentType, "application/json") {
				err = json.Unmarshal(requestBody, requestPost)
				if err != nil {
					fmt.Println("Executing State Method", pResponse, err)

				} else {
					// It manege to catch the name of the machine(HostName)
					hostname, err := os.Hostname()
					if err != nil {
						// In case does not catch the Host name
						log.Error("It has occur an Error when the HostName", err)
					}
					fmt.Println("Executing State Method POST inside")
					state := StateService{"ok", "Tuesday", "6:42 AM", hostname, "ruote"}
					js, _ := json.Marshal(state)
					pResponse.Header().Set("Content-Type", "application/json")
					pResponse.Write(js)
				}

			}
		}

	})
}
