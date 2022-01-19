package main

import (
	"log"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/subschedule"
)

var STRIPE_KEY = "sk_test_..."
var CUSTOMER_ID = "cus_..."
var FULL_PRICE = stripe.String("price_...")
var ZERO_PRICE = stripe.String("price_...")
var SCHEDULED_SUBSCRIPTION_ID = "sub_sched_..."

func main() {
	stripe.Key = STRIPE_KEY
	createNewFiveYearSubscription()
	//shortenSubscriptionByEighteenMonths()
}

func createNewFiveYearSubscription() {
	params := &stripe.SubscriptionScheduleParams{
		Customer:    stripe.String(CUSTOMER_ID),
		EndBehavior: stripe.String("cancel"),
		StartDate:   stripe.Int64(1641953064), // 11-jan-2022
		DefaultSettings: &stripe.SubscriptionScheduleDefaultSettingsParams{
			CollectionMethod: stripe.String("charge_automatically"),
		},
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this first price is the full $100 amount
						Price:    FULL_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1673489064), //11-jan-2023
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			// the four prices below represent the 4 years of $0 1-year subscriptions
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1705025064), //11-jan-2024
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1736647464), //11-jan-2025
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1768183464), //11-jan-2026
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1799719464), //11-jan-2027
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
		},
	}
	ss, err := subschedule.New(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Subscription successfully scheduled: %+v", ss)
}

func shortenSubscriptionByEighteenMonths() {
	subID := SCHEDULED_SUBSCRIPTION_ID // get this ID from the terminal output of createNewFiveYearSubscription()
	// the function calls here are essentially similar/identical to creating a brand new subscription schedule,
	// except we have fewer scheduled phases here, and the end date of the fourth phase is
	// 11-july-2025
	params := &stripe.SubscriptionScheduleParams{
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this first price is the full $100 amount
						Price:    FULL_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				StartDate:         stripe.Int64(1641953064), // 11-jan-2022
				EndDate:           stripe.Int64(1673489064), //11-jan-2023
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			// the four prices below represent the 4 years of $0 1-year subscriptions
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1705025064), //11-jan-2024
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1736647464), //11-jan-2025
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						// this second price is the $0 subscription
						Price:    ZERO_PRICE,
						Quantity: stripe.Int64(1),
					},
				},
				EndDate:           stripe.Int64(1752282264), //11-jul-2025
				ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorNone)),
			},
		},
	}
	ss, err := subschedule.Update(subID, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Subscription successfully updated: %+v", ss)
}
