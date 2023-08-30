# The CRUD Golang project

With this project you create a server which works with database.
There are several endpoint available

POST: "/book/" - Create book \
GET: "/book/" - Get books \
GET: "/book/{id}" - Get book \
PUT: "/book/{id}" - Update book \
DELETE: "/book/{id}" - Delete book \

You will need to create ".env" file in the root directory where you need to provide
ypur MySQL user name and the password

USER="root" \
PASSWORD="qwerty" \
URL="127.0.0.1" \
PORT=":3306"

# Add required packages
go mod tidy

# Go to the right path
cd ./cmd/main

# Build the project
go build

# Start using it
./main