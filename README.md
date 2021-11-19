# Account server
This is the back-end service for the [Account](https://github.com/writethesky/account) application
## Depend server
* [Basic GRPC Server](https://github.com/writethesky/basic)

## For developer
You can use the commands in the makefile

* install tools: `make install-tools`
* generate files necessary for the program to run: `make generate`
* run server: `make run`

Other commands you might be interested in

* clear protocol file: `make clean`
* pull the protocol file`make fetch-proto`

### View the API documentation

1. `make generate`
2. use your browser to visit `http://localhost:8080/swagger/index.html`

