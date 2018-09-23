# db-loadtest
> simply utility to load test postgres or mysql

This utility is designed to provide basic timing results for queries dispatcged against a database.  The utility is designed to execute files that are read from a file; this can be done sequentially or randomly. The utility can be run for a specified duration or until the queries file has been exhausted.  If run for a specified duration and the queries file has been exhausted, we will loop through the file again until the specified time has elapsed.  Queires can be run at a specified rate per second or with a fixed concurrency.


### Options

| Option | Description                           | Required | Default |
| ------ |-------------------------------------- | -------- | --------
| f      | path to queries file                  |     true |         |
| cs     | Connection string                     |     true |         |
| rps    | Requests Per Second                   |    false |         |
| c      | Concurrency (ignored if rps is set)   |    false |       1 |
| d      | Duration in seconds                   |    false |         |
| eof    | End test when query file is exhausted |    false |    true |
| s      | Process file sequentially             |    false |    true |
| r      | Process file randomly                 |    false |   false |


### Example Usage:

Two Request Per Second, sequentially from queries.json file, against pg, for 100 seconds
```
./db-loadtest -rps=2 -f=queries.json -cs=postgres://usr:pwd@host:5432/db -d=100
```

Two Request Per Second, randomly from queries.json file, against pg, for 100 seconds
```
./db-loadtest -rps=2 -f=queries.json -r -cs=postgres://usr:pwd@host:5432/db -d=100
```

Two Request Per Second, sequentially from queries.json file, against pg until query file is exhausted
```
./db-loadtest -rps=2 -f=queries.json -eof -cs=postgres://usr:pwd@host:5432/db
```

Five concurrent queries (using 5 connections), sequentially from queries.json file, against pg until query file is exhausted
```
./db-loadtest -c=5 -f=queries.json -eof -cs=postgres://usr:pwd@host:5432/db
```

### Example Output

```
Index: 1, Time: 22 ms, Error: None
Index: 0, Time: 100 ms, Error: None
Index: 3, Time: 26 ms, Error: None
Index: 2, Time: 34 ms, Error: None
Index: 5, Time: 37 ms, Error: None
Index: 4, Time: 37 ms, Error: None
Index: 6, Time: 24 ms, Error: None
Index: 7, Time: 24 ms, Error: None
Index: 8, Time: 25 ms, Error: None
Index: 9, Time: 27 ms, Error: None
Index: 10, Time: 24 ms, Error: None
Index: 11, Time: 28 ms, Error: None
```