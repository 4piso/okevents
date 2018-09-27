# okevents
just a proof of concept of a full cloud computer software

Events services

3 routes for events
    * Find Event
    * load all events
    * Create new Event

We're using mongodb for quickly approach, but we're planning to add other DB

# Persistence Layer
We add a small persistence layer for the database and all the interface that use data, the persistence layer is inside the persistence package which is only data level. The data is on state and the rest api will only use it to responde to the client
here is the definition [**Persistence Layer**](https://en.wikipedia.org/wiki/Persistence_(computer_science)
