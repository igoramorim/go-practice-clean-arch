# What is it?

This is a simple Go project aimed to apply a clean arch structure.

It is divided as the following:

- /api - http files to help while testing the app
- /cmd - a place to hold the main files
    - /cli - root command
    - /order - start the order system
- /config - loads the configurations
- /internal - all the internal packages
    - adapters - everything that "plugs" into the core app business rules
    - application - all the use cases implementations
    - domain - all the core app business packages and its rules
- migration - database migrations scripts
- pkg - all the packages the app exports and may be used by other apps

## How to run

Setup the env variables
```
make env
```

Build the CLI app. This will create a _/bin_ folder with the CLI executable
```
make build-cli
```

Setup the database and queue running the _docker-compose.yaml_ file
```
docker-compose up -d
```

Execute the order system
```
make order-system
```

### Usages

#### REST
Use the _/api/order.http_ file to create and list orders.

#### gRPC
Start the _evans_ client.
```
make grpc-evans
```
Set the package and service. Then execute the calls.
```
package pb
service OrderService
call CreateOrder
...
call FindAllOrdersByPage
```

#### GraphQL
Access _http://localhost:8080/playground_ and execute the queries or mutations.

#### CLI
Execute the root command to get help
```
./bin/cli order
```
And the sub commands [create] or [list] to see how to use them.
```
./bin/cli order create
...
./bin/cli order list
```
