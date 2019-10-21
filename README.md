# Test

## Run

**make up** - to run in docker-compose

other cmds you can find in makefile


## Routes

#### Create User

POST "/user/create"

{"userId":1, "transactionId":501, "type":"Win", "amount":50.5, "token":"testtask"}



#### Get User information
POST "/user/get"



#### Add Deposit
POST "/user/deposit"

{"userId":1, "depositId":1, "amount":50, "token":"testtask"}



#### Add Transaction
POST "/transaction"

{"userId":1, "transactionId":1, "type":"Win", "amount":50.5, "token":"testtask"}
