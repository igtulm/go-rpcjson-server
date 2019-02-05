## JSON-RPC Server

### Start

To start the server just run:

```
docker-compose up
```

Server is exposed to a port 1234.

### Usage:

Create account:
```
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "user.Create", "params": [{"Login":"NEW_LOGIN_IS_HERE"}]}' http://127.0.0.1:1234/rpc
```

Update account login:
```
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "user.Update", "params": [{"Login":"OLD_LOGIN_IS_HERE", "NewLogin":"NEW_LOGIN_IS_HERE"}]}' http://127.0.0.1:1234/rpc
```

Get account by login:
```
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "user.GetByLogin", "params": [{"Login":"CURRENT_LOGIN_IS_HERE"}]}' http://127.0.0.1:1234/rpc
```

### Testing:

For testing run:
```
docker-compose run --rm app make test
```
