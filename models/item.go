package models

type Items struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Rating          int    `json:"rating"`
	Category        string `json:"category"`
	Image           string `json:"image"`
	Reputation      int    `json:"reputation"`
	Price           int    `json:"price"`
	Availability    int    `json:"availability"`
	ReputationBadge string `json:"reputation_badge"`
}
