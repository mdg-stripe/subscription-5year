# Scheduled subscription demo
This simple demo shows how to create a subscription in Stripe with a 5 year duration. This is accomplished by charging the full amount in year 1; and then scheduling 4 1-year phases at $0 per year after year 1.

The demo also shows how the duration of the subscription can be shortened.

## Demo instructions
Prerequisites: install the Go compiler (https://go.dev/dl/).

1. Create a Stripe account, and log into the Dashboard.
2. Create a Customer with a saved payment method (use one of the test cards here: https://stripe.com/docs/testing#cards).
3. Create a Product.
4. In the Product object, create two prices: one with a price of $100, recurring yearly; and one with a price of $0, recurring yearly.
5. In `main.go`, update the first four `var` values with your Stripe secret key, the customer ID, and the two price ID's.
6. Open a terminal and type: `go run main.go`. Verify that the subscription was successfully created in the Stripe Dashboard. Note the subscription end date of January 2027.
7. From the output in the terminal, copy the ID value, and paste it into the `var SCHEDULED_SUBSCRIPTION_ID` parameter in `main.go`.
8. Comment out `createNewFiveYearSubscription`, and uncomment `shortenSubscriptionByEighteenMonths`.
9. In the terminal, type `go run main.go` again.
10. Verify that the subscription end date has been changed to July 2025.