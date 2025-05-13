#  Product API - Golang + Echo

A RESTful Product API built with Go and the Echo framework, showcasing standard backend development practices. This API supports full CRUD operations. 
---

##  Features

* Full RESTful API for managing products
* CRUD operations:

  * Create, Retrieve, Update (PUT/PATCH), and Delete products
* Logging middleware integration using **logrus**

---

##  Tech Stack

* **Language**: Go
* **Framework**: Echo
* **Database**: MongoDB
* **Logging**: [logrus](https://github.com/sirupsen/logrus)

---

## Endpoints

| Method | Endpoint                | Description                             |
| ------ | ----------------------- | --------------------------------------- |
| POST   | `/products`             | Create a new product                    |
| GET    | `/products`             | List all products                       |
| GET    | `/products/:pdoruct_id` | Retrieve a specific product             |
| PUT    | `/products/:pdoruct_id` | Update an existing product              |
| PATCH  | `/products/:pdoruct_id` | Partially update a product              |
| DELETE | `/products/:pdoruct_id` | Delete a product                        |
| GET    | `/search`               | According to query params list products |

---

## Search Endpoint Features

* Search by **partial or exact** product name
* **Filter** by price
* **Sort** by price (ascending or descending)

---
## Usage

### Example Request Body For 'POST /products'
```bash
{
    "Name":         "string",
    "Price":        float,
    "Description":  "string"
}
```
---
#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "data": "Object Id"
  }
}
```
---
### Example Request For 'GET /products'
The endpoint has pagination as default limit=10, page=1. if you want to as page=**param**, limit=**param**, you can change.

---
#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "data": [
      {
        "ProductId": "Id",
        "Name": "name",
        "Price": price,
        "Description": "description",
        "Created_at": "date"
      },
      {
        "ProductId": "Id",
        "Name": "name",
        "Price": price,
        "Description": "description",
        "Created_at": "date"
      },
      {
        .
        .
        .
      },
    ]
  }
}
```
---
### Example Request For 'GET /products/:product_id'

---
#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "data": {
      "ProductId": "Id",
      "Name": "name",
      "Price": price,
      "Description": "description",
      "Created_at": "date"
    }
  }
}
```
---
### Example Request Body For 'PATCH /products/:product_id'
```bash
{
    "Name":   "string",
}
```
---
#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "data": "Object Id"
  }
}
```
---

### Example Request Body For 'PUT /products/:product_id'
```bash
{
    "Name":         "string",
    "Price":        float,
    "Description":  "string"
}
```
---
#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "data": "Object Id"
  }
}
```
---
### Example Request For 'DELETE /products/:product_id'
```bash
{
    "Name":         "string",
    "Price":        float,
    "Description":  "string"
}
```
---
#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "data": "Succesfuly delete the product"
  }
}
```
---

### Example Request For 'GET /search'
The endpoint has default limit=10, page=1
* GET /search?name=**param**&price_min=**param**&price_max=**param**&page=**param**&limit=**param**&sort=**asc or desc** 
* GET /search?name=**param**&price_min=**param**
* GET /search?name=**param**&price_max=**param**
* GET /search?name=**param**&price_min=**param**&price_max=**param**
* GET /search?name=**param**&price_min=**param**&price_max=**param**&page=**param**
* GET /search?name=**param**&price_min=**param**&price_max=**param**&page=**param**&limit=**param**
* GET /search?price_min=**param**&price_max=**param**&page=**param**&limit=**param**&sort=**asc or desc** 

#### _Response Format_
```bash
{
  "stauscode": 200,
  "message": true,
  "data": {
    "products": [
      {
        "ProductId": "Id",
        "Name": "name",
        "Price": price,
        "Description": "description",
        "Created_at": "date"
      },
      {
        "ProductId": "Id",
        "Name": "name",
        "Price": price,
        "Description": "description",
        "Created_at": "date"
      },
      {
        .
        .
        .
      },
    ]
  }
}
```
---