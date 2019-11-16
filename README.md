# go-test-rest-api
a test api written in golang can be distributed by docker

to compile:
go build

to run:
	Windows and Mac Os:
		go-test-rest-api
	Linux:
		./go-test-rest-api

to build docker image: 
	Windows and Mac Os:
		docker build -t go-tester-rest-api .
	Linux:
		sudo docker build -t go-tester-rest-api .

to run docker image:
	Windows and Mac Os:
		docker run -p 4000:4000 go-tester-rest-api
	Linux:
		sudo docker run -p 4000:4000 go-tester-rest-api

to find container name:
	Windows and Mac Os:
		docker ps
	Linux: 
		sudo docker ps

to stop docker image:
	Windows and Mac Os:
		docker stop <container name>
	Linux:
		sudo docker stop <container name> 


