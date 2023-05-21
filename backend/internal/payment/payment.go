package payment

import (
	"fmt"
	"os"

	"github.com/osscameroon/jobsika/internal/graphql"
)

type IClient interface {
	CreateTier() error
	DeleteTier() error
}

const TierDescription = `
You are about to pay %.2f EUR for a job offer on jobsika.com.

Please proceed to payment and once done, you will receive an email with the job offer details.
`

const OSSCAMEROON_SLUG = "osscameroon"
const OSSCAMEROON_ID = "4rxg0j35-lzkwm6vz-bxbqvoe9-8n47daby"
const OPEN_COLLECTIVE_CONTRIBUTE = "https://opencollective.com/osscameroon/contribute/"

// PostTierResponse
type PostTierResponse struct {
	CreateTier struct {
		ID          string `json:"id"`
		LegacyID    int64  `json:"legacyId"`
		Slug        string `json:"slug"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Amount      struct {
			Value        float64 `json:"value"`
			Currency     string  `json:"currency"`
			ValueInCents int     `json:"valueInCents"`
		} `json:"amount"`
		Button interface{} `json:"button"`
		Goal   struct {
			Value        interface{} `json:"value"`
			Currency     string      `json:"currency"`
			ValueInCents interface{} `json:"valueInCents"`
		} `json:"goal"`
		Type              string      `json:"type"`
		Interval          string      `json:"interval"`
		Frequency         string      `json:"frequency"`
		Presets           interface{} `json:"presets"`
		MaxQuantity       int         `json:"maxQuantity"`
		AvailableQuantity int         `json:"availableQuantity"`
		CustomFields      interface{} `json:"customFields"`
		AmountType        string      `json:"amountType"`
		MinimumAmount     struct {
			Value        interface{} `json:"value"`
			Currency     string      `json:"currency"`
			ValueInCents interface{} `json:"valueInCents"`
		} `json:"minimumAmount"`
		EndsAt            interface{} `json:"endsAt"`
		InvoiceTemplate   string      `json:"invoiceTemplate"`
		UseStandalonePage bool        `json:"useStandalonePage"`
		SingleTicket      bool        `json:"singleTicket"`
	} `json:"createTier"`
}

//create a struct
type Client struct {
	graphQLClient *graphql.Client
}

type OpenCollectiveOptions struct {
	URL string
	KEY string
}

func NewClient(opts OpenCollectiveOptions) (*Client, error) {
	client := graphql.NewClient(opts.URL, opts.KEY)
	if client == nil {
		return nil, fmt.Errorf("failed to create graphql client")
	}

	return &Client{
		graphQLClient: client,
	}, nil
}

func (c Client) CreateTier() (PostTierResponse, error) {
	query := graphql.Query(`
	mutation (
  	  $tier: TierCreateInput!
  	  $account: AccountReferenceInput!
	) {
  	  createTier(tier: $tier, account: $account) {
    	id
    	legacyId
    	slug
    	name
    	description
   	   amount {
      	  value
      	  currency
      	  valueInCents
    	}
    	button
    	goal {
      	  value
      	  currency
      	  valueInCents
    	}
    	type
    	interval
    	frequency
    	presets
    	maxQuantity
    	availableQuantity
    	customFields
    	amountType
    	minimumAmount {
      	  value
      	  currency
      	  valueInCents
    	}
    	endsAt
    	invoiceTemplate
    	useStandalonePage
    	singleTicket
  	  }
	}
`)

	price := 5.0
	name := "jobsika-joboffer"
	if os.Getenv("ENVIRONMENT") == "development" {
		price = 0.01
		name = "test-jobsika-joboffer"
	}
	variables := map[string]interface{}{
		"tier": map[string]interface{}{
			"amount": map[string]interface{}{
				"value":    price,
				"currency": "EUR",
			},
			"name":              name,
			"description":       fmt.Sprintf(TierDescription, price),
			"button":            "PAY NOW",
			"type":              "PRODUCT",
			"amountType":        "FIXED",
			"frequency":         "ONETIME",
			"maxQuantity":       1,
			"useStandalonePage": true,
		},
		"account": map[string]interface{}{
			"id":   OSSCAMEROON_ID,
			"slug": OSSCAMEROON_SLUG,
		},
	}

	response := PostTierResponse{}

	if err := c.graphQLClient.Run(query, variables, &response); err != nil {
		return PostTierResponse{}, err
	}

	return response, nil
}
