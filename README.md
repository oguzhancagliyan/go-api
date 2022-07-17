# Prerequisites

Docker have to be installed


# How to run

Go to the root file and execute the command

> docker-compose up


# Fundamentals

The API has 3 endpoints which are
- Get the value of the given key.
- Set the value of the key
- Flush database

The API backups data to src/backups/TIMESTAMP-data.json every minute.

At every cold start, the API, reads the files from src/backups/ and collect the data which are saved into files before and insert the current database.


# Endpoints

- Get the value of the given key

    - Endpoint is http://localhost:8080/api/oguzhan
    - Http method GET

The endpoint can return different variation

|         Key       |Response                          |StatusCode                         |
|----------------|-------------------------------|-----------------------------|
|oguzhan|{"error": {"message": "key not found"}}            | 404|
|oguzhan          |{"data": {"key": "oguzhan","value": "cagliyan"}}|200|
|          |{"error": {"message": "key shouldn't be empty!"}}|400|




- Set the value of the key
    - Endpoint is http://localhost:8080/api/
    - Http method POST
    - Request body contains 2 parameters
        - Key : string
        - Value : string

The endpoint can return different variation :

|         Body       |Response                          |StatusCode                         |
|----------------|-------------------------------|-----------------------------|
|{"Key" : "oguzhan","Value" : "cagliyan"}|{"data": {"key": "oguzhan","value": "cagliyan"}}| 201 |
|{"Key" : "" ,"Value" : "cagliyan"}          |{"error": {"message": "key shouldn't be empty!"}}|400|
|  {"Key" : "oguzhan","Value" : "cagliyan"}        |{"error": {"message": "key already exist"}}|409|


* Flush the database
    * Endpoint is http://localhost:8080/api/
    * Http method DELETE

| Body | Response |StatusCode                         |
|------|----------|-----------------------------|
|      |          | 204 |


- Postman collection

You can find in the root folder with "collection.json" name




## Backup

An endless loop provide the backup functionality. We provide the functionality with go routines. Every minute endless loop runs and backup files to src/backups/TIMESTAMP-data.json


## Restore Db

At every cold start API reads the backups folder and collect data after that insert to the Db. 
