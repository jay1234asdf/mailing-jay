package model

import "time"

type Details struct {
	Recipient      string    `json:"recipient,omitempty" validate:"required"`
	Amount         string    `json:"amount,omitempty" validate:"required"`
	Date           time.Time `json:"date,omitempty" `
	Status         string    `json:"status,omitempty" validate:"required,eq=pending|eq=completed|eq=failed"`
	Notice         string    `json:"notice,omitempty" validate:"required"`
	Tax            string    `json:"tax,omitempty" validate:"required"`
	Total          string    `json:"total,omitempty"`
	Method         string    `json:"method,omitempty" validate:"required"`
	Country_County string    `json:"country_county,omitempty"`
}
