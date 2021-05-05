Test task for Avito trinee/junior developer https://github.com/avito-tech/adv-backend-trainee-assignment

## ENDPOINTS

### 1. Create 
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

### 2. Single Ad
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