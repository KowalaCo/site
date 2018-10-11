# site
A hacktoberfest project, The kowala.co api and frontend

# API docs
## /api/v1/user/create (POST)
### Request Data
{"username": "TheUsername", "password": "ThePassword"}
### Normal Response
{"token": "TheToken"}

## /api/v1/user/authenticate (POST)
### Request Data
{"username": "TheUsername", "password": "ThePassword"}
### Normal Response
{"token": "TheToken"}

## /api/v1/user/me (GET)
### Request Data
Headers: {"Authorization": "TheToken"}
### Normal Response
{"id": "TheId", "username": "TheUsername", "role": 0, "verified": false}