# wallet

The goal is to write a JSON API in Golang to get the balance and manage credit/debit operations on
the players wallets. For example, you might receive calls on your API to get the balance of the
wallet with id 123, or to credit the wallet with id 456 by 10.00 €. The storage mechanism to use will
be MySQL.
 
## Usage
```
make local
```

## Pre-inserted data
Wallets with IDs 123 (balance 0.01) and 456 (balance 12.30)

## Endpoints

- [X] balance : retrieves the balance of a given wallet id  
GET `http://localhost:5001/api/v1/wallets/{wallet_id}/balance` 
### Output example 
```
[
    {
        "balance": "1"
    }
]
```
- [X] credit : credits money on a given wallet id  
POST `http://localhost:5001/api/v1/wallets/{wallet_id}/credit`
### Input example 
```
[
    {
        "balance": "1"
    }
]
```
### Output example 
```
[
    {
        "balance": "13.3",
        "message": "ok"
    }
]
```
- [X] debit : debits money from a given wallet id  
POST `http://localhost:5001/api/v1/wallets/{wallet_id}/debit`
### Input example 
```
[
    {
        "balance": "1"
    }
]
```
### Output example 
```
[
    {
        "balance": "11.3",
        "message": "ok"
    }
]
```

## Business rules
- [X] A wallet balance cannot go below 0.
- [X] Amounts sent in the credit and debit operations cannot be negative.

## Bonus
- [X] Cache the wallet balances in Redis, so that they can be fetched from cache
- [ ] Add auth endpoint and authentication verification
- [X] Add unit tests for the business rules/logic
- [X] Log the incoming requests

## Libraries to use
- HTTP : https://github.com/gin-gonic/gin
- MySQL : https://github.com/go-gorm/gorm
- Redis : https://github.com/go-redis/redis
- Numbers : https://github.com/shopspring/decimal
- Logger : https://github.com/sirupsen/logrus

## Notes
- No need to care about the currencies.
- No need to create wallets, they can be pre-populated in storage.
- Make sure to return some meaningful errors if an operation is not possible.
- A particular attention will be put on how the application is constructed, more
specifically how the web layer, repositories for storage, entities and business logic are structured.  
Ideally, this architecture should make the application testable, and not too dependent on implementation details (for example, which repository/storage mechanism we use etc...)
