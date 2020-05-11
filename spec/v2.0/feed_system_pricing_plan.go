package gbfsspec

type (
	FeedSystemPricingPlans struct {
		Metadata

		Data SystemPricingPlansData `json:"data"`
	}

	SystemPricingPlansData struct {
		Plans []SystemPricingPlan `json:"plans"`
	}

	SystemPricingPlan struct {
		// Identifier for a pricing plan in the system.
		PlanID string `json:"plan_id"`

		// URL where the customer can learn more about this pricing plan.
		URL string `json:"url,omitempty"`

		// Name of this pricing plan.
		Name string `json:"name"`

		// Currency used to pay the fare. (ISO 4217)
		Currency string `json:"currency"`

		// Currency used to pay the fare.
		Price Price `json:"price"`

		// Will additional tax be added to the base price?
		IsTaxable Boolean `json:"is_taxable"`

		// Customer-readable description of the pricing plan. This should include the duration, price,
		// conditions, etc. that the publisher would like users to see.
		Description string `json:"description"`
	}
)

func (_ FeedSystemPricingPlans) FeedKey() string {
	return FeedKeySystemPricingPlans
}
