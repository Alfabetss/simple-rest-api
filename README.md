# Simple-rest-api

create simple API using echo framework

## Getting Started
in this repository we will implement echo framework to create APIs that implement common http methods namely GET, POST, PUT, DELETE. 

### Prerequisites

some things that must have or be installed to run this program : 

* [golang](https://golang.org/doc/install) - The Programming Language
* [mysql](https://www.mysql.com/downloads/) - The Database
* [vscode](https://code.visualstudio.com/download) - The IDE
* [postman](https://www.postman.com/downloads/) - The API Client

#### Create Table
To create a table you can copy the script in the /script/*.sql folder and run in your database

#### Configuration 

| Name | Type | Default | Required | Example Value |
| --- | --- | --- | --- | --- |
| Port | String | :8888 | Yes | :8080
| DBName | String | - | Yes | example_database |
| DBProtocol | String | - | Yes | tcp |
| DBHost | String | - | Yes | localhost |
| DBPort | String | - | Yes | 3360 |
| DBUser | String | - | Yes | root |
| DBPass | String | - | Yes | administrator |

### Running
clone repository :
```
git clone https://github.com/Alfabetss/simple-rest-api.git
```
open terminal : 
```
cd path/to/program
```
run the program : 
```
go run server.go
```
open new terminal, and run : 
```
curl "http://localhost:1122/talent" \
  -X POST \
  -d "{\n  \"name\": \"robert jr\",\n  \"experience\": [\n    {\n      \"companyName\": \"marvel\"\n    },\n    {\n      \"companyName\": \"sony\"\n    }\n  ]\n}" \
  -H "Content-Type: application/json" 
```

## Built With

* [Echo](https://echo.labstack.com/) - The web framework used
* [squirrel](https://github.com/Masterminds/squirrel) - The Query Builder
* [viper](https://github.com/spf13/viper) - The Configuration tools
