# ethereum-go-scanner

This is a golang service to connect to ethereum mainnet and retrieve the blocks and store them in postgres.
This service provides an endpoint "/transactions/{AccountId}" that retrives the transactions done by the AccountID provided

### Steps to run the service
First you need to get your postgres database up and running and create a database and table in it to store the transactions.
This can be done using the docker-compose file given in the repo.
To run docker-compose file you need to have docker and docker-compose installed on your system.
You just need to run the following command in terminal `docker-compose run` in the project folder. This will run a postgres database and create the schema in a docker container.

Next you need to start the service that can be done by opening a terminal and navigating to the project folder. Run the command `go run main.go` which will run the service that can be accessed at `localhost:8010/`
