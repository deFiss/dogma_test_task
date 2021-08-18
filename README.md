Dogma test task
========================
RESTful API for storing, adding, modifying and deleting user information.   

## Project architecture
It was based on Uncle Bob's article about "[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)"  
It provides the following benefits:
>- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
>- Independent of UI. The UI can change easily, without changing the rest of the system.
>- Independent of Database. Your business rules are not bound to the database.
>- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

#### This project does not fully disclose all the above advantages of the architecture. This is just a demonstration of the possibilities. 
***
## Deployment
```sh
# Clone the repository
$ git clone git@github.com:deFiss/dogma_test_task.git
$ cd dogma_test_task

# Create .env file, change variables if needed 
$ cp .env.example .env

# Start all services
$ make run
```
Next, you need to migrate the database  

_One could try to create a database every time the project starts, but this is a bad approach in the long run_

For migration, the cli tool [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate) is used, you need to install it on the system according to the instructions from the repository.

```sh
# Migrate database
$ make migrate
```

## Usage
You can use a swagger for testing.  
After launching the project, it will be available at http://127.0.0.1:8000/swagger/index.html 

***
## Dependencies
- **github.com/gin-gonic/gin** - Web server
- **github.com/go-openapi/swag** - API documentation
- **github.com/jmoiron/sqlx** - SQL rapid driver
- **github.com/joho/godotenv** - Loading environment variables from .env file
- **github.com/lib/pq** - Postgresql driver