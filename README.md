# goji-api-sample
Simple C-R-U-D APIs

# Run
```
$ go run main.go
```

## Create
```
$ curl -X GET 'http://localhost:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]

$ curl -X POST -d '{ "title": "title", "article": "article" }' 'http://localhost:9999/'
Report Created! id: 4

$ curl -X GET 'http://localhost:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900},{"id":4,"title":"title","article":"article","created_at":1470562930,"updated_at":1470562930}]
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
$ curl -X GET 'http://localhost:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]

$ curl -X PUT -d '{ "title": "TITLE", "article": "ARTICLE" }' 'http://localhost:9999/3'
Report Updated! id: 3

$ curl -X GET 'http://localhost:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"TITLE","article":"ARTICLE","created_at":1470495800,"updated_at":1470564112}]
```

## Delete
```
$ curl -X GET 'http://localhost:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]

$ curl -X DELETE 'http://localhost:9999/1'

$ curl -X GET 'http://localhost:9999/'
[{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]
```
