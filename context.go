package main

import (
	"time"
)

// Offer is a data structure that represents a published offer.
type Offer struct {
	// The datetime at which the offer was registered in the site.
	CreationDate time.Time `json:"date"`
	// The position that the job offer is valid for.
	Position string `json:"position"`
	// The company that published this offer.
	Company string `json:"company"`
	// The list of tags that are related to this offer.
	Tags []string `json:"tags"`
	// The content of the offer description.
	Description string `json:"description"`
	// The URL that identifies this offer.
	URL string `json:"url"`
	// The unique ID that identifies this offer.
	ID int `json:"id"`
}

// Context holds the application state.
type Context struct {
	location Location
	offers   []Offer
	idIndex  map[int]Offer
	tagIndex map[string][]int
}

// updateIndices destroys and re-creates the id index and the tag index
// associated with a context. These are useful for filtering offers and for
// using offer IDs around in the application.
func (context *Context) updateIndices() {
	context.idIndex = make(map[int]Offer)
	context.tagIndex = make(map[string][]int)

	// presenceIndex tries to make faster guessing which tags are in tagIndex.
	presenceIndex := make(map[string]bool)
	for _, offer := range context.offers {
		// Put the offer in the ID index.
		context.idIndex[offer.ID] = offer

		// Extract tags and put the offer in the tag index.
		for _, tag := range offer.Tags {
			if !presenceIndex[tag] {
				// This tag was never added to the tagIndex in first place.
				context.tagIndex[tag] = make([]int, 0)
				presenceIndex[tag] = true
			}

			// Put this offer in the index associated with the tag.
			context.tagIndex[tag] = append(context.tagIndex[tag], offer.ID)
		}
	}
}

// SetOffers manually set the list of offers and updates the indices.
func (context *Context) SetOffers(offers []Offer) {
	context.offers = offers
	context.updateIndices()
}

// SetOffersByLocation retrieves the offers in the given location, and
// places them in the current context. It also will update the offers
// index used by the application.
func (context *Context) SetOffersByLocation(location Location) error {
	offers, err := FetchOffers(location)
	if err != nil {
		return err
	}
	context.location = location
	context.SetOffers(offers)
	return nil
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
