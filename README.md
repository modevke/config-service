# Rest API with swagger

## Installation process
* Install swagg ``go get -u github.com/swaggo/swag/cmd/swag``
* Install all dependancies ``go mod download``
* Create .env file in the root of your project and copy from .env.example
* You will first need to generate the swagger doc by running the command ``swag init``
* Run the main.go file to start the server
* You can access the swagger documentation on `http://localhost:8030/api/documentation/index.html`