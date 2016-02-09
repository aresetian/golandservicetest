package main

import (
	  "net/http"
	  
	)

func main{

	  // Se inicia el servidor
 	  err := http.ListenAndServe(":"+cfgConfig.ConfigHotetec.PortServer, nil)
	  // Si se produce un error con el listener de la aplicacion, se imprimi- el error
	  if err != nil {
	    // si se produce un error al iniciar el servidor se imprime en el log
	    log.Error("El servicio no pudo ser iniciado", err.Error())
	
	  }
}
