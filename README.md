# backend-collection-go-1

### Version: release-1.0 (2022-12-18)

### Summary:
This is Go backend number 1 from my backend collection project. This backend is a REST API for CRUD student data build with Gin Gonic framework, MySQL, and GORM, also tested with Apitest framework.

### Requirements:
1. go (tested: v1.19.3 windows/amd64, v1.18.4 windows/amd64)
2. mysql (tested: v8.0.31)

### Steps to run the backend server:
1. install all requirements
2. clone repository `https://github.com/reyhanfikridz/backend-collection-go-1` at directory `$GOPATH/src/github.com/reyhanfikridz/`
3. at repository root directory, which is `$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-1`:
    1. switch to branch release-1.0 with `git checkout release-1.0`
    2. download required go modules with `go mod download`
    3. copy all downloaded go modules to the repository with `go mod vendor`
    4. create file .env with contents:

    ```
    MYSQL_DBUSER="<mysql database user, example:mysql>"
    MYSQL_DBPASS="<mysql database password>"
    MYSQL_DBNAME="<mysql database name, example:backend_collection_go_1>"
    MYSQL_DBTESTNAME="<mysql database test name, example:backend_collection_go_1_test>"
    MYSQL_DBHOST="<mysql database host, example:localhost>"
    MYSQL_DBPORT="<mysql database port, example:3306>"

    GIN_HOST="<app (gin) host, example:127.0.0.1>"
    GIN_PORT="<app (gin) port, example:3000>"
    ```

    5. create mysql databases with name same as in .env file
4. at `$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-1/cmd/migration`, migrate database with `go build` then `./migration`
5. at repository root directory again, which is `$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-1`, test server first with `go test ./...`
6. at `$GOPATH/src/github.com/reyhanfikridz/backend-collection-go-1/cmd/backend_collection_go_1`, run server with `go build` then `./backend_collection_go_1`

### API collection:
1. Go to https://www.postman.com/reyhanfikri/workspace/backend-collection-go-1/overview
2. Choose `release-1.0` collection

### License:
This project is MIT license, so basically you can use it for personal or commercial use as long as the original LICENSE.md included in your project.
