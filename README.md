# Simple CRUD Backend with Golang

## How to run this project

```scheme
1. Clone this repository to $GOPATH/src/github.com/joedha8/
2. Create "animestoredb" in Postgres
3. Import "animestoredb.sql" in files directory

"psql animestoredb < animestoredb.sql"

4. Run `./AnimeStore` to run the server
5. Then, open browser and go to "http://localhost:1212"
```

## RESTful API

### Category

Delete  : `http://localhost:1212/category/{id}`

Get All : `http://localhost:1212/category`

Get One : `http://localhost:1212/category/{id}`

Insert  : `http://localhost:1212/category`

With Body

```scheme
{
	"category_name": "Aksesoris",
	"created_by": "b6d1589b-d068-48e0-9e77-265482acfb98"
}
```

Update  : `http://localhost:1212/category/{id}`

With Body

```scheme
{
	"category_name": "Accessoris"
}
```

### Detail Order

Delete  : `http://localhost:1212/detail_order/{id}`

Get All : `http://localhost:1212/detail_order`

Get One : `http://localhost:1212/detail_order/{id}`

Insert  : `http://localhost:1212/detail_order`

With Body

```scheme
{
	"id_order": "7a5bb0ea-67ba-4316-a544-1a81327bad17",
	"product_id": "7a5bb0ea-67ba-4316-a544-1a81327bad17",
	"quantity": 9,
	"sub_total": 8,
	"created_by": "b6d1589b-d068-48e0-9e77-265482acfb98"
}
```

Update  : `http://localhost:1212/detail_order/{id}`

With Body

```scheme
{
	"quantity": 9,
	"sub_total": 9000
}
```

### Order

Delete  : `http://localhost:1212/order/{id}`

Get All : `http://localhost:1212/order`

Get One : `http://localhost:1212/order/{id}`

Insert  : `http://localhost:1212/order`

With Body

```scheme
{
	"customer_id": "7a5bb0ea-67ba-4316-a544-1a81327bad17",
	"total_price": 190000,
	"created_by": "b6d1589b-d068-48e0-9e77-265482acfb98"
}
```

Update  : `http://localhost:1212/order/{id}`

With Body

```scheme
{
	"total_price": 250000
}
```

### Product

Delete  : `http://localhost:1212/product/{id}`

Get All : `http://localhost:1212/product`

Get One : `http://localhost:1212/product/{id}`

Insert  : `http://localhost:1212/product`

With Body

```scheme
{
	"id_category": "7a5bb0ea-67ba-4316-a544-1a81327bad17",
	"description": "Gundam 00 Diver Ori Bandai",
	"price": 190000,
	"stock": 8,
	"product_name": "Gundam 00 Diver",
	"created_by": "b6d1589b-d068-48e0-9e77-265482acfb98"
}
```

Update  : `http://localhost:1212/product/{id}`

With Body

```scheme
{
	"description": "Jual Dakimakura Murah",
	"price": 250000,
	"stock": 9,
	"product_name": "Dakimakura"
}
```

### Wishlist

Delete  : `http://localhost:1212/wishlist/{id}`

Get All : `http://localhost:1212/wishlist`

Get One : `http://localhost:1212/wishlist/{id}`

Insert  : `http://localhost:1212/wishlist`

With Body

```scheme
{
	"customer_id": "7a5bb0ea-67ba-4316-a544-1a81327bad17",
	"product_id": "7a5bb0ea-67ba-4316-a544-1a81327bad17",
	"created_by": "b6d1589b-d068-48e0-9e77-265482acfb98"
}
```

Update  : `http://localhost:1212/wishlist/{id}`

With Body

```scheme
{
	"customer_id": "7a5bb0ea-67ba-4316-a544-1a81327bad17h",
	"product_id": "7a5bb0ea-67ba-4316-a544-1a81327bad17"
}
```