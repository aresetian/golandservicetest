#!/bin/bash


#Config Logs
export LOG_PATH=${LOG_PATH:-/logs/}
export LOG_NAME=${LOG_NAME:-Hotetec_$HOSTNAME.log}
export LOG_LEVEL=${LOG_LEVEL:-1} 
#LevelDebug = 1,        LevelInfo  = 2, LevelWarn  = 3, LevelError = 4



echo "Generating config.ini..."
envsubst < /.setup/config.ini > /go/src/golandservicetest/golangservicetest/config.ini


# Start dataservice
cd /go/src/golandservicetest/golangservicetest/ && ./golangservicetest -fileconfig=config.ini
