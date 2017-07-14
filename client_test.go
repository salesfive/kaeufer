package kaeufer

import "github.com/salesfive/kaeufer"

// Example without config
func Example() {
	c := kaeufer.Client{
		Id:     "client id",
		Secret: "client secret",
	}

	leads, err := c.Leads()
}

// Example with optional config
func Example() {
	c := kaeufer.Client{
		Id:     "client id",
		Secret: "client secret",
	}

	leads, err := c.Leads(
		kaeufer.PerPage(10),
		kaeufer.FromTimestamp(1459157310),
	)
}
