package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/salesfive/kaeufer"
)

func main() {
	// demo account from the api documentation.
	clientId := "9971557793fdf19f4bf96de5d48be46547bdc5ec"
	clientSecret := "f5dd4d8943a51447ef63b061ccd0aab43e98e1f0"

	c := kaeufer.Client{
		ID:     clientId,
		Secret: clientSecret,
	}

	leads, err := c.Leads(
		kaeufer.PerPage(1),
		kaeufer.FromTimestamp(1459157309),
	)

	spew.Dump(leads)
}
