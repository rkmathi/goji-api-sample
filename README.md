# goji-api-sample
Simple C-R-U-D APIs

# Run
```
$ go run main.go
```

## Create
```
```

## Read (Index)
```
$ curl -X GET 'http://localhost:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]
```

## Read (Show)
```
$ curl -X GET 'http://localhost:9999/1'
{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600}

$ curl -X GET 'http://localhost:9999/0'
Report Not Found
```

## Update
```
```

## Delete
```
```
