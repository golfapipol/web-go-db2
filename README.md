# Example project (Golang + Nginx + DB2)

## Run a Demo

```bash
docker-compose up -d
```

## Access via Web

check it at [http://localhost:3000](http://localhost:3000)

## Access via API

### Get All Data 

```
curl http://localhost:4000/employee
```

example response

```json
[
    {"firstname":"CHRISTINE","lastname":"HAAS","gender":"F"},
    {"firstname":"MICHAEL","lastname":"THOMPSON","gender":"M"},
    {"firstname":"SALLY","lastname":"KWAN","gender":"F"},
    {"firstname":"JOHN","lastname":"GEYER","gender":"M"},
    {"firstname":"IRVING","lastname":"STERN","gender":"M"},
    {"firstname":"EVA","lastname":"PULASKI","gender":"F"},
    {"firstname":"HELENA","lastname":"WONG","gender":"F"},
    {"firstname":"ROY","lastname":"ALONZO","gender":"M"}
]
```

### Transform Data 

```
curl --location --request POST 'http://localhost:4000/employee/transform' --header 'Content-Type: application/json' --data '{"firstname":"EVA","lastname":"PULASKI","gender":"F"}'
```

example response

```json
{"firstname":"EVA","gender":"F","genderCode":2,"lastname":"PULASKI"}
```

## Run Unit Test

```bash
cd api && go test ./business
```