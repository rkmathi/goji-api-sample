# goji-api-sample
Simple CRUD APIs

# Run
```
$ go get goji.io

$ go run main.go
```

## Read (Index)
```
$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]
```

## Read (Show)
```
$ curl -X GET 'http://127.0.0.1:9999/1'
{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600}

$ curl -X GET 'http://127.0.0.1:9999/0'
ERROR! Report not found.
```

## Create
```
$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]

$ curl -X POST -d '{ "title": "title", "article": "article" }' 'http://127.0.0.1:9999/'
New report id: 4

$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900},{"id":4,"title":"たいとる","article":"あーてぃくる","created_at":1470791481,"updated_at":1470791481}]
```

## Update
```
$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]

$ curl -X PUT -d '{ "title": "TITLE", "article": "ARTICLE" }' 'http://127.0.0.1:9999/3'
Update report id: 3

$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"たいとる","article":"あーてぃくる","created_at":1470495800,"updated_at":1470791563}]
```

## Delete
```
$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":1,"title":"title1","article":"this is article1","created_at":1470495600,"updated_at":1470495600},{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]

$ curl -X DELETE 'http://127.0.0.1:9999/1'

$ curl -X GET 'http://127.0.0.1:9999/'
[{"id":3,"title":"title3","article":"here is article3","created_at":1470495800,"updated_at":1470495900}]
```
