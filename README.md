### Boiler Plate

Boiler plate REST API using [gorilla/mux v1.8.0](https://github.com/gorilla/mux).

##### Call Parameter
```sh
# using curl
curl -X GET -H "Content-Type: application/json" -d '{}' \
http://localhost:8081/
```

##### Example ENV Config
```env
APPLICATION_NAME=privy
APPLICATION_PORT=8081
APPLICATION_TRACE_REPORT=false

DATABASE_ENABLE=true
DATABASE_HOST=139.59.123.79
DATABASE_PORT=3306
DATABASE_USERNAME=root
DATABASE_PASSWORD=P@ssw0rd
DATABASE_NAME=faltar
```