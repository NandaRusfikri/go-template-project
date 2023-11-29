# Backend API Template

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.20
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)


## Swagger 
http://localhost:54321/swagger/index.html

## Install Dependency The Project

To start working, execute

```
go get .
```

## Running

To run the project, run the following command:

```
go run main.go
```

You should be able to access the API at http://localhost:54321


## Running With Docker

To run the project, run the following command:

```
docker-compose up -d
```


## Generate Swagger Doc

To run the project, run the following command:

```
make swag
```

You should be able to access the API at http://localhost:54321/swagger/index.html


## Folder Structure Pattern

```
│── constant
│   └── constant.go
├── docs
│   └── swagger.json
│   └── swagger.yaml
└── database
│   └── migrate_sql.go
│   └── mysql.go
│   └── postgres.go
├── pkg
│   └── gin.go
│   └── jwt.go
│   └── smtp.go
└── docker
│   └── swagger
│   │     └── Dockerfile
│   │     └── openapi.yml
│   └── mysql
│   │     └── Dockerfile
│   │     └── mysql.cnf
│   └── golang
│   │     └── Dockerfile
├── module
│   └── auth
│   │     └── entity
│   │     │     └── auth.go
│   │     └── repository
│   │     │     └── auth.go 
│   │     └── usecase
│   │     │     └── auth.go
│   │     └── controller
│   │     │     └── auth.go  
│   │     └── interface.go
│   └── user
│   │     └── entity
│   │     │     └── user.go
│   │     └── repository
│   │     │     └── user.go 
│   │     └── usecase
│   │     │     └── user.go
│   │     └── controller
│   │     │     └── user.go  
│   │     └── interface.go
└── presentation
│   └── server
│   │     └── server.go
└── dto
│   └── auth.go
│   └── user.go
└── util
│   └── helper.go
│   └── random.go

```

## Folder Status And Description

- #### Entity

| **Folder Name** | **Description**                                                                               |
|-----------------|-----------------------------------------------------------------------------------------------|
| *entity*        | *Kumpulan struct yang menggambarkan struktur data di table database atau dari data external.* | 


- #### Controller

| **Folder Name** | **Description**                                                                                                                                                                                                                                                                               |
|-----------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| *controller*    | *Kumpulan fungsi yang digunakan untuk menghandle semua request yang di kirimkan dari client melalui routing, dimana nanti requestnya tersebut akan di teruskan ke layer usecase dan layer repository untuk di proses lebih lanjut, yang nanti nya akan digunakan untuk aplikasi itu sendiri.* | 

- #### Repositorys

| **Folder Name** | **Description**                                                                                                    |
|-----------------|--------------------------------------------------------------------------------------------------------------------|
| *Repositorys*   | *Kumpulan fungsi yang digunakan untuk mengambil atau mengirim data. sumber nya bisa dari database, Rest API  dll * | 

- #### Use Case

| **Folder Name** | **Description**                            |
|-----------------|--------------------------------------------|
| *usecase*       | *Kumpulan fungsi yang berisi logic bisnis* | 

- #### DTO (Data Transfer Object)
| **Folder Name** | **Description**                                                                                                                                                                                              |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| *dto*           | *Kumpulan Object yang digunakan untuk merepresentasikan structure request yang di inginkan, sesuai dengan request yang di perlukan oleh database, yang nanti nya akan digunakan untuk aplikasi itu sendiri.* | 
