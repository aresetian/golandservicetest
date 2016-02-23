package main

import (
	"gopkg.in/gcfg.v1"
	"net/http"
	"golandservicetest/golangservicetest/handler"
	"golandservicetest/golangservicetest/log"
	"golandservicetest/golangservicetest/util"
	"flag"
)


// This var is used to save the name file of the configuracion file
var fileConfig = flag.String("fileconfig", "", "The configuration file")


func initLogs(pPath, pNameFile string, pLogLevel int) {

   log.InitLog(pPath+pNameFile, pLogLevel)
}

func main(){


      // It is used to get the name of the file
	  flag.Parse()
	
	  var cfgConfig util.Config
	
	
	  // It reads the configuration file and it will filled the fileConfig Struct
	  err1 := gcfg.ReadFileInto(&cfgConfig, *fileConfig)
	  if err1 != nil {
	    panic(err1)
	  }
	
	  // It sets up the Log File
	  initLogs(cfgConfig.Logs.LogPath, cfgConfig.Logs.LogName, cfgConfig.Logs.LogLevel)
      //initLogs("/home/ubuntu/workspace/logs/", "DockerServiceLog.log", 1)	


      //This method contains the mapping between URL and Handlers.
	  handler.SetUpHandlers("c9")
	  
	  // If debug is enable, this message will be printed
	  log.Debug("Server Started")
	  
	  // The server is started
	  // Only the ports 8080, 8081 y 8082  can be visibles on the Internet when it works with cloud9
 	  err := http.ListenAndServe(":8080", nil)
	  
	  // If an Error occur
	  if err != nil {
	    log.Error("The service could be started", err.Error())
      }

}
