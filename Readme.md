This is a basic template that configures a Golang Webserver restAPI authentication using sessions. It uses GORM to connect the Go code to mySQL.

notes for prod https:
1. change routes.Setup() to include extra session Config.
2. Change CORS AllowOrigins to only client ip