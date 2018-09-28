# okevents
just a proof of concept of a full cloud computer software

Events services

3 routes for events
- Find Event
- load all events
- Create new Event

We're using mongodb for quickly approach, but we're planning to add other DB

# Persistence Layer
We add a small persistence layer for the database and all the interface that use data, the persistence layer is inside the persistence package which is only data level. The data is on state and the rest api will only use it to responde to the client
here is the definition [**Persistence Layer**](https://en.wikipedia.org/wiki/Persistence_(computer_science)

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
