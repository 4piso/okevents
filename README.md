# okevents 
just a proof of concept of a full cloud computer software, using queue systems

Events services

3 routes for events
- Find Event
- load all events
- Create new Event

We're using mongodb for quickly approach, but we're planning to add other DB

# Persistence Layer
We add a small persistence layer for the database and all the interface that use data, the persistence layer is inside the persistence package which is only data level. The data is on state and the rest api will only use it to responde to the client
here is the definition [**Persistence Layer**](https://en.wikipedia.org/wiki/Persistence_(computer_science))

# Rest API for Event Service
The Rest API is a simple request/response to the persistence layer, there is to main request the user can do:
- `/events` - This will get all the events on the database
- `/events/id/12 or /events/name/cea - get the event base on name or id`

For the rest api we're using gorilla tool kit, and using the mux for url variables as a query
here is an example of the json format to create new events, we're using 3 simple documents, name, location and halls
```
{
    "name": "best",
    "startdate": 768346784368,
    "enddate": 768346784368,
    "duration": 120,
    "location": {
        "name": "Paris fifa park",
        "address": "le monses",
        "country": "frn",
        "opentime": 10,
        "closetime": 15,
        "halls": [{
            "name": "lobby",
            "location": "principal building",
            "capacity" : 150
        }]
    }
}
```

The system will save the data on the POST /events route. 

# Asynchronous Layer

We're trying to implement the Publish/Subscribe pattern communication instead of client request/response. Each Publisher can emit message to the Subscriber who will get the messages.
We're using the Rabbit MQ Advanced Message Queueing Protocol to show how this is posible on a large scale project. the documentation for the library that we're using is on the following link: [**RabbitMQ for GO**](https://godoc.org/github.com/streadway/amqp)

# Docker for Rabbit MQ
$ docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
You will need to have docker installed on your machine,to start docker you will need the start command line
$ docker run rabbitmq
Then you can go to http://localhost:15672 and you will see the rabbit MQ administrator open

# Booking Services
We're adding a new services, the Booking Services. for now is using the same mongo db layer than the other event services project, also using the same msgqueue package, for a large scale project we will need to separate this approach

# Run the Apps
You need to run the 2 apps, to see the example. Events Services will create a services and the booking services will print the same events on the command line, for now. on the future ticket this will save the data to the database


