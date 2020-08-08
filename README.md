# proomea helper library for golang services

## MISSING

-

## Minor Updates
Version  | Date | Summary
------------- | ------------- | -------------
v1 | 07.08.2020 | init documentation

## Description and Functionality
Contains 

- HTTP Rest-Errors
- JWT Token Handling


#### gRPC endpoints
No | Endpoint  | Type | Description| JWT | PERMISSION 
------------- |------------- | ------------- | ------------- | --------------  | -------------- 
1 | /api/device/types  | GET | Get all device Types | YES | SUPERUSER, ADMIN, DEFAULT

	
##### Customer
Gives back all devices types which are selectable

###### Schema











```
curl --location --request GET '46.101.179.215:8000/api/device/get/602d0c19-cbfd-11ea-a978-00ffb1b25530' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.
eyJpYXQiOjE1OTQ2NDM4NjEsIm5iZiI6MTU5NDY0Mzg2MSwianRpIjoiZjIxZmNiODMtZmQ3OC
00YmYwLThhODQtOTBiZjJiOTk5OGM2IiwiZXhwIjoxNTk0NjQ0NzYxLCJpZGVudGl0eSI6IjQx
MTAwZjNmNmJiYTRkYTJiMGU1ZGQ5M2EwNWI1N2NlIiwiZnJlc2giOnRydWUsInR5cGUiOiJhY2
Nlc3MiLCJ1c2VyX2NsYWltcyI6eyJ1c2VySWQiOiI0MTEwMGYzZjZiYmE0ZGEyYjBlNWRkOTNh
MDViNTdjZSIsInJvbGUiOjIsIm9yZ2FJZCI6IjE4NmQ3ZjhmNDcyZTQyZWZiOTZhYmVkNTdkZG
FiNmUwIn19.mUInDcz82p9TflmtI7nXIWxI3_pkT0n0O_x8xFx5Vfg'
```
###### Response example
```
