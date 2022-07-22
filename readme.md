File Parser App
======================

Application is created, for parsing large port file and inserting in database.
Data is read in chunks and inserted in inbuild database (map)
Data is logged after insertion from database (map)

Data is placed in resources > ports.json file

Starting the application
---------------------

Place the .json file at resources > ports.json and fire command:
```
$ docker-compose up
```

Test the application
---------------------
```
$ go test ./test/*.go
```

Note
---------------------
Application is created in 1 hour 
Needs further work on: 
   - MongoDB could be used for data insertion in database
   - Better package distribution