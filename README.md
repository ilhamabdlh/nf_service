# NFT Service

NFT Service is a RESTful API for managing non-fungible tokens (NFTs). It allows users to create, retrieve, update, and delete NFT items, as well as purchase them.

## Installation and Setup

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/nft-service.git
   
2. Navigate to the project directory
   ```bash
   cd nft_service

3. Create a config.json file in the project root and add the following environment variables with your database credentials:
 ```bash
{
    "database": {
      "host": "localhost",
      "port": "3306",
      "user": "root",
      "password": "",
      "name": "nft_service"
    }
}
```

4. Create Table with the query:
```bash
CREATE TABLE items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    rating INT NOT NULL,
    category VARCHAR(50) NOT NULL,
    image VARCHAR(255) NOT NULL,
    reputation INT NOT NULL,
    price INT NOT NULL,
    availability INT NOT NULL,
    reputation_color VARCHAR(255) NOT NULL,
);
```

## Running
1. Build and run the application:
```bash
go build -o nft-service
./nft-service
```

2. The API server should now be running locally. You can access it at http://localhost:8080

## API Endpoints
```bash
GET /items: Retrieve all items.
GET /items/:id: Retrieve an item by ID.
POST /items: Create a new item.
PUT /items/:id: Update an existing item.
DELETE /items/:id: Delete an item.
POST /items/:id/purchase: Purchase an item.
```

3. Testing API with this schema:
```bash
1. Retrieve all items:
Method: GET
Endpoint: http://localhost:8080/api/items

2. Retrieve an item by ID:
Method: GET
Endpoint: http://localhost:8080/api/items/:id
Replace :id with the actual ID of the item you want to retrieve.

3. Create a new item:
Method: POST
Endpoint: http://localhost:8080/api/items
Body:
{
    "name": "New Item One",
    "rating": 3,
    "category": "sketch",
    "image": "https://example.com/image.jpg",
    "reputation": 600,
    "price": 30,
    "availability": 20
}

4. Update an existing item:
Method: PUT
Endpoint: http://localhost:8080/api/items/:id
Replace :id with the actual ID of the item you want to update.
Body:
{
    "name": "Updated Item One",
    "rating": 4,
    "category": "photo",
    "image": "https://example.com/updated_image.jpg",
    "reputation": 700,
    "price": 40,
    "availability": 15
}

5. Delete an item:
Method: DELETE
Endpoint: http://localhost:8080/api/items/:id
Replace :id with the actual ID of the item you want to delete.

6. Purchase an item:
Method: POST
Endpoint: http://localhost:8080/api/purchase/:id
Replace :id with the actual ID of the item you want to purchase.

```



## Notes
- Ensure that your MySQL database is running and accessible before starting the API server.
- Replace placeholders in the config.json file with your actual database credentials.
