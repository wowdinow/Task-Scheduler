# TASK SCHEDULER

This task scheduler is an application to manage scheduled task

- RESTful endpoint for asset's CRD operation
- JSON formatted response

## GET /task

_Request Header_

```
not needed
```

_Request Body_

```
not needed
```

_Response (200 - Ok)_

```
[
    {
        "Execution Time": "2023-11-10T17:13:00+07:00",
        "Task ID": 1,
        "Task Name": ""
    },
    {
        "Execution Time": "2023-11-10T17:13:00+07:00",
        "Task ID": 2,
        "Task Name": ""
    },
    {
        "Execution Time": "2023-11-10T17:13:00+07:00",
        "Task ID": 3,
        "Task Name": ""
    }
]
```

## POST /task

_Request Header_

```
not needed
```

_Request Body_

```
{
    "cron" : "<cron time format>",
    "task_name" : "<task name>",
    "exec" : "<execution cmd>"
}
```

_Response (200 - Ok)_

```
{ "message" : "Success add task with ID : " + id}
```

## DELETE /task/:id

_Request Header_

```
not needed
```

_Request Body_

```
not needed
```

_Request Params_

```
{ id : <id>}
```

_Response (200 - Ok)_

```
{ "message": `Delete Task with ID : ` + id }
```

## GLOBAL ERROR

_Response (404 - Not Found)_

```
404 - Not Found
```

_Response (500 - Internal Server Error)_

```
500 - Internal Server Error
```