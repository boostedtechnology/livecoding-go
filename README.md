# Boosted Live Coding

We are creating a simple ledger API using Go, specifically with the GORM and Gorilla Mux libraries.


1. Clone the repository
2. Run `go mod tidy` to download the dependencies
3. Run `go run main.go` to start the server


## Task 1: Implement DELETE /accounts/:id endpoint

We want to be able to delete an account by its ID. Add a new endpoint to the controller to delete an account by its ID.
Note, we are using "soft deletes" in this application, so the account will not be deleted from the database, but will be marked as deleted.

See reference: https://gorm.io/docs/delete.html

Run `go test tests/account_test.go` to see if your new endpoint is working! You may need to update the test depending on
the implementation.


## Task 2: Implement POST /transactions/

**Context**

A transaction in double-entry accounting is a record of a financial transaction. Each transaction has a description and a list of entries. Each entry has an account ID, a debit amount, and a credit amount. Within a transaction, the sum of the debit amounts must equal the sum of the credit amounts. There must be at least two entries in a transaction.

**Task**

Create the service, controller, and routes for the POST /transactions endpoint.

* Only worry about the POST endpoint for now. The other CRUD operations do not need to be implemented.

* The models are already created for Transaction and Entry. You may want or need to add GORM tags to improve performance (e.g. indexes, foreign keys, etc. - https://gorm.io/docs/indexes.html).

You can decide how to structure the request body, but the following is potential example that shows a transaction with three account entries on the transaction.

```json
{
  "description": "Office supplies at Staples",
  "entries": [
    {
      "accountId": "a018f63f-3794-4927-90fa-b62f26892203",
      "type": "DEBIT",
      "amount": 12456
    },
    {
      "accountId": "ee51e463-f4f9-4ddf-a2b3-af8e196f851f",
      "type": "CREDIT",
      "amount": 10000
    },
    {
      "accountId": "d9e645c9-ff24-4360-9cde-c31fcffa76dc",
      "type": "CREDIT",
      "amount": 2456
    }
  ]
}
```

## Task 3: Implement GET /accounts/:id/transactions/

Retrieve all the transactions for a given account.


## Task 4: Add a "balance" field to the response of GET /accounts/:id

**Context**
Each account has a "balance" based on the transactions involving the account.

* For Asset and Expense accounts, the balance is the sum of all the debit entries involving the account minus the sum of all the credit entries involving the account.
* For Liability, Equity, and Revenue accounts, the balance is the sum of all the credit entries involving the account minus the sum of all the debit entries involving the account.

**Task**

On the response of GET /accounts/:id, add a "balance" field on the response.


## Task 5: Prevent deletion of accounts with non-zero balance

Update the DELETE /accounts/:id endpoint to prevent the deletion of accounts with non-zero balance.

Make sure to include a test for the updated logic.
