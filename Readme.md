This is a basic template that configures a Golang Webserver restAPI authentication using sessions. It uses GORM to connect the Go code to mySQL.

notes for prod https:
1. change routes.Setup() to include extra session Config.
2. Change CORS AllowOrigins to only client ip

backlog:

primary todo:
1. refactor file structure
2. documentation

secondary todo:
1. change AllowOrigins: "*", to localhost or client server
2. make sure middleware allows access of certain routes without a session
3. encrypt .env