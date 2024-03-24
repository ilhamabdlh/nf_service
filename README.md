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



Installation and Setup
1. lone the repository to your local machine:
git clone https://github.com/yourusername/nft-service.git

2. Navigate to the project directory:
cd nft-service

3. Create a config.json file in the project root and add the following environment variables with your database credentials:
{
    "database": {
      "host": "localhost",
      "port": "3306",
      "user": "root",
      "password": "",
      "name": "nft_service"
    }
}
  

Running
1. Build and run the application
go build -o nft-service
./nft-service

2. The API server should now be running locally. You can access it at 
http://localhost:8080

API Endpoints
GET /items: Retrieve all items.
GET /items/:id: Retrieve an item by ID.
POST /items: Create a new item.
PUT /items/:id: Update an existing item.
DELETE /items/:id: Delete an item.
POST /items/:id/purchase: Purchase an item.

Notes
Ensure that your MySQL database is running and accessible before starting the API server.
Replace your_database_host, your_database_port, your_database_username, your_database_password, and your_database_name in the .env file with your actual database credentials.

