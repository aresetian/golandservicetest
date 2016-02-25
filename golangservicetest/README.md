# golandservicetest


How to run the service on Cloud9
First, Going to 
/home/ubuntu/workspace/src/golandservicetest/golangservicetest
Second, Execute the command
./golandservicetest -fileconfig=config.ini
Finally,
Going to the URL
https://golang-test-service-carlos-andres.c9users.io/c9/State
https://localhost:8080/c9/State(Check if there is diferrence between https and http on Docker)
Where
c9 is the context and
State is the URI that was config in the project



I created the docker image with 
docker build -t service-go-andres .

I run a container with 
docker run  -d -p 4020:8080  service

http://localhost:4020/c9/State

