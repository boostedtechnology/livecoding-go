# Boosted Live Coding

We are creating a simple ledger API using Go, specifically with the GORM library for database operations.


1. Clone the repository
2. Run `go mod tidy` to download the dependencies
3. Run `go run main.go` to start the server


## Task 1: Implement DELETE /accounts/:id/ endpoint

We want to be able to delete an account by its ID. Add a new endpoint to the controller to delete an account by its ID. Upon successful deletion, return 204 No Content.

GORM uses "soft deletes" by default, so the account will not be deleted from the database, but will be marked as deleted. This is likely not relevant to the task, but just something to note.

See reference: https://gorm.io/docs/delete.html

Run `go test tests/account_test.go` to see if your new endpoint is working! You may need to update the test depending on the implementation.


## Task 2: Implement POST /transactions/

**Context**

A transaction in double-entry accounting is a record of a financial transaction. Each transaction has a description and a list of entries. Each entry has an account ID, an amount, and a type (DEBIT or CREDIT). See below for an example data model.

Within a transaction,

* the sum of the debit amounts must equal the sum of the credit amounts

* no account should be used in more than one entry

* there must be a non-zero amount of entries for a transaction

**Task**

Create the service, controller, and routes for the POST /transactions/ endpoint. Upon successful creation, return 204 No Content.

* Only worry about the POST endpoint. The other CRUD operations do not need to be implemented.

* The models are already created for Transaction and Entry. You may want or need to add GORM tags to improve performance (e.g. indexes, foreign keys, etc. - https://gorm.io/docs/indexes.html).

You can decide how to structure the request body, but the following is a potential example that shows a valid transaction with three entries.

```json
{
  "description": "Office supplies at Staples",
  "entries": [
    {
      "accountId": 1,
      "type": "DEBIT",
      "amount": 12456
    },
    {
      "accountId": 45,
      "type": "CREDIT",
      "amount": 10000
    },
    {
      "accountId": 23,
      "type": "CREDIT",
      "amount": 2456
    }
  ]
}
```


## Task 3: Add a "balance" field to the response of GET /accounts/:id/

**Context**

Each account has a "balance" based on the transactions involving the account.

* For Asset and Expense accounts, the balance is the sum of all the debit entries involving the account minus the sum of all the credit entries involving the account.
* For Liability, Equity, and Revenue accounts, the balance is the sum of all the credit entries involving the account minus the sum of all the debit entries involving the account.

**Task**

On the response of GET /accounts/:id, add a "balance" field on the response.


## Task 4: Implement GET /accounts/:id/transactions/

Retrieve all the transactions for a given account.

The response format is up to you to decide.


## Task 5: Prevent deletion of accounts involved in transactions

Update the DELETE /accounts/:id/ endpoint to prevent the deletion of accounts involved in any transaction.

Make sure to include a test for the updated logic.
