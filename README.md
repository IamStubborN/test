# Test

Architecture inspired from https://github.com/bxcodec/go-clean-arch

Have much overcoding, but i can easily add new features and unit testing for everything

## Run

**make up** - to run in docker-compose

other cmds you can find in makefile

folder **httptest** for testing routes via jetbrains http client

## Routes

#### Create User

POST "/user/create"

{"id":1, "balance":0.0, "token":"testtask"}



#### Get User information
POST "/user/get"

{"id":1, "token":"testtask"}



#### Add Deposit
POST "/user/deposit"

{"userId":1, "depositId":1, "amount":50, "token":"testtask"}



#### Add Transaction
POST "/transaction"

{"userId":1, "transactionId":1, "type":"Win", "amount":50.5, "token":"testtask"}
