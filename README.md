# Boosted Live Coding

We are creating a simple ledger API using Go, specifically with the GORM and Gorilla Mux libraries.

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to download the dependencies
3. Run `go run main.go` to start the server

## Improvements

#### Step 1: Implement DELETE /accounts/:id endpoint

We want to be able to delete an account by its ID. Add a new endpoint to the controller to delete an account by its ID.
Note, we are using "soft deletes" in this application, so the account will not be deleted from the database, but will be marked as deleted.

See reference: https://gorm.io/docs/delete.html

Make sure to include a test for the new endpoint.


#### Step 2: Implement POST /transactions

A transaction in double entry account is a record of a financial transaction that affects at least two accounts. Each transaction has a description and a list of entries. Each entry has an account ID, a debit amount, and a credit amount.

The models are already created, but the "service" layer, controller, and routes are missing. Use the /accounts endpoint
as a guide to implement the POST /transactions endpoint.

Make sure to include a test for the new endpoint.


#### Step 3: Implement GET /accounts/:id/transactions

Retrieve all the transactions for a given account.


#### Step 4: Add a "balance" field to the response of GET /accounts/:id

For Asset accounts, the balance is the sum of all the debit entries involving the account minus the sum of all the credit entries involving the account. For Liability and Equity accounts, the balance is the sum of all the credit entries
involving the account minus the sum of all the debit entries involving the account. In short, the "natural" state of an
asset account is to have a debit balance, and the "natural" state of a liability or equity account is to have a credit balance.

Make sure to include a test for the new endpoint.


#### Step 5: Prevent deletion of accounts with non-zero balance

Update the DELETE /accounts/:id endpoint to prevent the deletion of accounts with non-zero balance. You can choose whether
to enforce this rule at the database level or at the application level.

Make sure to include a test for the updated logic.
