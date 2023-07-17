# taskStore

### Prerequisite
1. Make sure your local development setup has go runtime installed. 
For installation, refer to [instructions](https://go.dev/doc/install).

### Running Web Application

Download dependancies and run the web application.
```
make all
```

### Testing Web Application

Run the tests.
```
make test
```

### Sample API invocations

1. List existing tasks.
```
curl --location 'http://0.0.0.0:8083/tasks' | jq .
[
  {
    "id": 1,
    "title": "P0",
    "content": "This is a high priority task"
  },
  {
    "id": 2,
    "title": "P2",
    "content": "This is a low priority task"
  }
]
```

2. Get details of specific task by taskId.

```
curl --location 'http://0.0.0.0:8083/tasks/1' | jq .
{
  "id": 1,
  "title": "P0",
  "content": "This is a high priority task"
}
```

3. Create a new task.

a. Simple request.
```
curl --location 'http://0.0.0.0:8083/tasks' --data '{
  "id": 3,
  "title": "P1",
  "content": "This is a medium priority task"
}' | jq .
{
  "id": 5,
  "title": "P1",
  "content": "This is a medium priority task"
}
```

b. Task create request with untrusted user generated data.

```
curl --location 'http://10.13.106.157:8083/tasks' \
--data '{
  "id": 3,
  "title": "P1 <a onblur=\"alert(secret)\" href=\"http://www.google.com\">Google</a>",
  "content": "This is a medium priority task"
}' | jq .
{
    "id": 3,
    "title": "P1 <a href=\"http://www.google.com\" rel=\"nofollow\">Google</a>",
    "content": "This is a medium priority task"
}

```

4. Delete existing task.

```
curl -v --location --request DELETE 'http://0.0.0.0:8083/tasks/3'
*   Trying 0.0.0.0...
* TCP_NODELAY set
* Connected to 0.0.0.0 (0.0.0.0) port 8083 (#0)
> DELETE /tasks/3 HTTP/1.1
> Host: 0.0.0.0:8083
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 204 No Content
< Content-Type: application/json
< Date: Mon, 17 Jul 2023 15:13:02 GMT
<
* Connection #0 to host 0.0.0.0 left intact
```

# Swagger UI

*[Note: Swagger definition files are present in `docs` folder.]*

![Screenshot 2023-07-17 at 9 40 26 PM](https://github.com/vinayakshnd/taskStore/assets/4953411/f53e052b-93ac-49f5-bd9f-0141bcbf0a69)
