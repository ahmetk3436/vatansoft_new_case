#!/bin/bash

# initiate replica set
mongo <<EOF
var config = {
    "_id": "dbrs",
    "version": 1,
    "members": [
        {
            "_id": 1,
            "host": "mongo1:27017",
            "priority": 3
        },
        {
            "_id": 2,
            "host": "mongo2:27017",
            "priority": 2
        },
        {
            "_id": 3,
            "host": "mongo3:27017",
            "priority": 1
        }
    ]
};
rs.initiate(config, { force: true });
rs.status();
EOF

# enable authentication
mongo admin <<EOF
use admin
db.createUser({user:"root", pwd:"root", roles:[{role:"root", db:"admin"}]})
EOF

# create stajdb database
mongo -u admin -p password <<EOF
use logdb
db.createCollection("logs")
EOF

