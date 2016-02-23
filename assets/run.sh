#!/bin/bash


#Config Logs
export LOG_PATH=${LOG_PATH:-/logs/}
export LOG_NAME=${LOG_NAME:-ServiceTest_$HOSTNAME.log}
export LOG_LEVEL=${LOG_LEVEL:-1} 
#LevelDebug = 1,        LevelInfo  = 2, LevelWarn  = 3, LevelError = 4



echo "Generating config.ini..."
envsubst < /.setup/config.ini > /go/src/golangservicetest/golangservicetest/config.ini


# Start Service
cd /go/src/golangservicetest/golangservicetest/ && ./golangservicetest -fileconfig=config.ini
