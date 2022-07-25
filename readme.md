File Parser App
======================

Application is created, for parsing large port file and inserting in database.
Data is read in chunks and inserted in inbuild database (map)
Data is logged after insertion from database (map)

Starting the application
---------------------

Need to create a tmp directory on local host system 
```
$ mkdir /tmp/test
```

Place the .json files in the new created directory.
```
$ cp SRC_PATH/ports.json /tmp/test
```

Place the .json file at resources > ports.json and fire command:
```
$ docker-compose up
```

Output shall be found on console and in /tmp/test/app.log file.

Note
---------------------
Application is created in 1 hour 
Needs further work on: 
   - MongoDB could be used for data insertion in database
   - Better logging features
