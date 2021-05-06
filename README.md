Test task for Avito trinee/junior developer https://github.com/avito-tech/adv-backend-trainee-assignment

## Description
It is a small web service written in Go for educational purposes. I only used stdlib (except db driver).

## How to start

### 1. Docker
1. Run command: `docker-compose up` \
service will run on localhost:8000

### 2. local
1. Run database: `docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres postgres:12` \
Thiw will run database
2. Run service: `go run main.go` \
This will run service and make migrations

## ENDPOINTS

### 1. Create ad
URL: `/create` - POST \
Payload:
```
{
    "name": "Vaz 2101",
    "price": 100.01,
    "description": "New Car",
    "photos": ["cdn://link.to.photo"]
}
```

Response:
```
{
    "id": 33,
    "name": "Vaz 2101",
    "description": "New Car",
    "price": 100.01,
    "photos": [
        "cdn://link.to.photo"
    ],
    "created": "2021-05-05T21:55:31.21025Z"
}
```

### 2. Get Ad
URL: `/get/{id}` - GET \
Params: You can pass "description" or/and "photos" with param "Fields" for more information \
`?fields=description&fields=photos` \
Full url: `http://localhost:8000/ad/8?fields=description&fields=photos`

Response:
```
{
    "id": 8,
    "name": "Vaz 2102",
    "price": 200.22,
    "photos": ["https://link-to-photo"],
    "description": "New car",
    "created": "2021-05-04T09:11:20.274002Z"
}
```


### 3. List ad
URL: `/asd` - GET \
Params: You can pass "page" for pagination \
Full url: `http://localhost:8000/ads?page=1`

Response:
```
[
    {
        "Id": 17,
        "Name": "Vaz 2102",
        "Price": 200.22,
        "Photos": ""
    }
]
```