# ⚡️ AxxonSoft Golang test task ⚡

## Table of contents
* 📖 [General info](#General info)
* 💻 [System requirements](#System requirements)
* 🌁 [Structure](#Structure)
* ⚙️ [Setup & Launch](#Setup&Launch)
* 📱 [Contacts](#Contacts)

## 📖 General info
Test task for Golang Developer

Write HTTP server for proxying HTTP-requests to 3rd-party services.
The server is waiting HTTP-request from client (curl, for example). In request's body there should be message in JSON format. For example:
```
{
    "method": "GET",
    "url": "http://google.com",
    "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
        ....
    }
}
```
Server forms valid HTTP-request to 3rd-party service with data from client's message and responses to client with JSON object:
```
{
    "id": <generated unique id>,
    "status": <HTTP status of 3rd-party service response>,
    "headers": {
        <headers array from 3rd-party service response>
    },
    "length": <content length of 3rd-party service response>
}
```
Server should have map to store requests from client and responses from 3rd-party service.

## 💻 System requirements
* Make
* Docker
* Git

## 🌁 Structure
``` text
├── cmd
│   └── ...
│  
├── config
│
├── internal
│   ├── api
│   │   └── ...
│   ├── middleware
│   │   └── ...
│   ├── server
│   │   └── ...
│   └── services
│       └── ...
├── pkg
│   ├── map
│   │   └── ...
│   ├── middleware
│   │   └── ...
│   ├── server
│   │   └── ...
│   ├── services
│   │   └── ...
│   └── utils
│       └── ...
│
├── common.env
├── Dockerfile
├── go.mod
├── Makefile
├── README.md
```

* [cmd/](cmd) - Application launch point is located in the directory
* [config/](config/) - Application configuration types, constansts, constructors located in the directory
* [internal/](internal/) - Directory contains specific implementations of interfaces located in [pkg/](pkg/),
  specific entities required for the application and several unit tests
* * [internal/api](internal/api) - Directory contains the incoming request handler and handler tests
* * [internal/middleware](internal/middleware) - Directory contains a specific wrapper for providing access to request handlers to a resource located in the directory [pkg/map](pkg/map)
* * [internal/server](internal/server) - Directory contains proxy server implementation
* * [internal/services](internal/services) - Directory contains specific route handlers and tests
* [pkg/](pkg/) - Directory contains general interfaces, specific resources implementation, util function and tests
* * [pkg/map](pkg/map) - Directory contains map to store client requests and responses
* * [pkg/middleware](pkg/middleware) - Directory contains a wrapper interface that provide public access to encapsulated specific resource
* * [pkg/server](pkg/server) - Directory contains server interface
* * [pkg/services](pkg/services) - Directory contains dto objects, util dto functions and tests
* * [pkg/utils](pkg/utils) - Directory contains application util functions
## ⚙️ Setup & Launch
``` bash
# 1) clone repo
git clone https://github.com/AlexScherba16/axxonsoft_golang_test_task

# 2) go to repo directory
cd axxonsoft_golang_test_task

# 3) build application
make build

# 4) run application
make run
```
While server is running you may send several requests like in example below
for more information, please check [Makefile](Makefile)
``` bash
make send_fail_request
# wait response

make ok_google_test
# wait response

make ok_httpbin_test
# wait response

# etc
```

## 📱 Contacts
``` 
email:      alexscherba16@gmail.com
telegram:   @Alex_Scherba
```
