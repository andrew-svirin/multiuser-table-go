# multiuser-table-go

Multiple users can edit one table at same time.

### Tech stack

Infrastructure: Docker  
Server: Go + Websocket + MySql + Javascript

## Client

- User can authorize on server
- User connecting by websocket to server
- User sending operations under the connection, table to server
- User listening to operations under the connection, table from server

# Server

- Serve api requests and websocket events in separate processes
- Avoid intersections in processes
- Listening for events from client and process them
- Sending events to other clients in parallel
- Storing in database state of table

### Development mode

1. Run `make start` to run containers
2. Run `make load-fixtures`to load fixtures
3. Run `make serve-server` to run server  
   ![img.png](/docs/files/server.jpg)
4. Open browser in `http://localhost:8080`
   ![app.gif](/docs/files/app.gif)