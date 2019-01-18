package main

import (
	"time"
)

// Offer is a data structure that represents a published offer.
type Offer struct {
	// The datetime at which the offer was registered in the site.
	CreationDate time.Time
	// The position that the job offer is valid for.
	Position string
	// The company that published this offer.
	Company string
	// The list of tags that are related to this offer.
	Tags []string
	// The content of the offer description.
	Description string
	// The URL that identifies this offer.
	URL string
	// The unique ID that identifies this offer.
	id int
}

// Context holds the application state.
type Context struct {
	offers   []Offer
	idIndex  map[int]Offer
	tagIndex map[string][]int
}

// NewContext will set up a new context object that represents these offers.
func NewContext(offers []Offer) *Context {
	// Create the context object.
	context := &Context{
		offers:   offers,
		idIndex:  make(map[int]Offer),
		tagIndex: make(map[string][]int),
	}

	// presenceIndex tries to make faster building the tag index.
	presenceIndex := make(map[string]bool)

	// Build the index.
	for _, offer := range offers {
		context.idIndex[offer.id] = offer

		for _, tag := range offer.Tags {
			if !presenceIndex[tag] {
				// Initialise new ID array for this new tag.
				context.tagIndex[tag] = make([]int, 0)
				presenceIndex[tag] = true
			}
			context.tagIndex[tag] = append(context.tagIndex[tag], offer.id)
		}
	}

	return context
}

func (c *Context) GetOffer(id int) *Offer {
	offer, ok := c.idIndex[id]
	if !ok {
		return nil
	}
	return &offer
}

func (c *Context) CountOffers() int {
	return len(c.offers)
}
