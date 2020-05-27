## A Video Processing API in GO

``go get github.com/saiprasadkrishnamurthy/interviews-api
``

## Swagger UI
``
go get -u github.com/swaggo/swag/cmd/swag
``

### 
``
../../../../bin/swag init --output public
``

The above must be run everytime you want to regenerate the swagger documentation.

SwaggerUI:
http://localhost:8083/

### Docker Build
```
docker build -t interviews-api .
```

### Docker Run
```
docker run -p 8083:8083 interviews-api
```


