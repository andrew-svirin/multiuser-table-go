# multiuser-table-go

Multiple users can edit one table at same time.

## Client

- User can authorize on server
- User connecting by websocket to server
- User sending operations to server

# Server

- Listening for signals from client and process them
- Sending signals to other clients in parallel
- Storing in database state of table in parallel
