Learn microservice app merchant use api gateway for access two serve is different languange.

That is for the Store service use Golang and Merchant Service use PHP.

For the Api Gateway use Golang.

# How to run

## Service Store
```
cd store
go run main.go
```

## Service Merchant
```
cd store
php artisan serve
```

## Service Api Gateway
```
cd store
go run main.go
```

# Route
## Merchant
Get data Merchant. For the access add header `Authorization` with value `merchant`.

```
http://localhost:8080/merchants
```

## Store
Get all data store. For the access add header `Authorization` with value `su-admin`.

```
http://localhost:8080/store
```
