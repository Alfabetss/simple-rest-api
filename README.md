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

create database mysql on your local machine : 
```
CREATE DATABASE resource;
```

create table talent & experience on your database
```
CREATE TABLE `talent` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
```

```
CREATE TABLE `experience` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `talent_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `talent_id` (`talent_id`),
  CONSTRAINT `experience_ibfk_1` FOREIGN KEY (`talent_id`) REFERENCES `talent` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
```

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
