# City Bikes

API client that consumes city bike (not citibike) data from the [City Bikes
API](https://api.citybik.es/v2/), caches it with Redis (using go-redis), and
measures the response time. 

Technology
----------
* Go
* Redis

Endpoints
---------

| Method     | URI                                  | Action                                      |
|------------|--------------------------------------|---------------------------------------------|
| `GET`      | `/api/networks`                      | `Retrieve all networks`                     |
| `GET`      | `/api/networks/{id}`                 | `Retrieve network`                          |


Sample Usage
---------------

`http get http://localhost:5000/api/networks`
```
    "networks": [
        {
            "id": "bbbike", 
            "location": {
                "city": "Bielsko-Biała", 
                "country": "PL", 
                "latitude": 49.8225, 
                "longitude": 19.044444
            }, 
            "name": "BBBike"
        }, 
        {
            "id": "bixi-montreal", 
            "location": {
                "city": "Montreal, QC", 
                "country": "CA", 
                "latitude": 45.5086699, 
                "longitude": -73.55399249999999
            }, 
            "name": "Bixi"
        }, 
...
```
logged to console after first request:  
`retrieved networks from remote api in 455.69015ms`

logged to console after second request:   
`retrieved networks from cache in 6.075998ms`

Run
---
With docker:
```
docker-compose build
docker-compose up
Go to http://localhost:5000 and visit one of the above endpoints
```

Alternatively, open `redis.go` and point the Redis URI to your server.

`cd` into `./city-bikes` (if you are not already); then run:
```
go build
./city-bikes
Go to http://localhost:5000 and visit one of the above endpoints
```
