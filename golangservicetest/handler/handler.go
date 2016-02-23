
package handler
	
import (
  "encoding/json"
  "net/http"
  "os"
  "golangservicetest/golangservicetest/log"
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
  http.HandleFunc("/"+pContext+"/State", State)
}


/*
   This struct contain the definition of the object that will be used on the respond.
*/
type StateService struct {
     Status string
  	 Day string	  
  	 Hour string
  	 HostName string
  	 LogPath string
}
	
	
/*

Description :
    Method that create a JSON Object and represents the data of an  StateService Struct.
Parameters :
    pResponse http.ResponseWriter :  It contains the response that it will be send to the final user.
    pRequest *http.Request :         It contains the request od the user.

Return :
       No data to return.

*/

func State(pResponse http.ResponseWriter, pRequest *http.Request) {
    
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
}
