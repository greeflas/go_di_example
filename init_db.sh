#!/usr/bin/env bash

rm -f example.db

sqlite3 example.db 'CREATE TABLE users(id INTEGER PRIMARY KEY ASC, name VARCHAR);'
sqlite3 example.db 'INSERT INTO users (name) VALUES ("John");'

sqlite3 example.db 'CREATE TABLE messages(id INTEGER PRIMARY KEY ASC, pattern VARCHAR);'
sqlite3 example.db 'INSERT INTO messages (pattern) VALUES ("Hello, %s!");'
