// Copyright (c) 2017 Philipp Weber
// Use of this source code is governed by the MIT license which can be found in the repositorys LICENSE file.

package hibpgo

import (
	"time"
)

// BreachModel contains a number of attributes describing the breach in detail.
type BreachModel struct {
	Name         string    `json:"Name"`
	Title        string    `json:"Title"`
	Domain       string    `json:"Domain"`
	BreachDate   string    `json:"BreachDate"`
	AddedDate    time.Time `json:"AddedDate"`
	ModifiedDate time.Time `json:"ModifiedDate"`
	PwnCount     int       `json:"PwnCount"`
	Description  string    `json:"Description"`
	DataClasses  []string  `json:"DataClasses"`
	IsVerified   bool      `json:"IsVerified"`
	IsFabricated bool      `json:"IsFabricated"`
	IsSensitive  bool      `json:"IsSensitive"`
	IsRetired    bool      `json:"IsRetired"`
	IsSpamList   bool      `json:"IsSpamList"`
	LogoType     string    `json:"LogoType"`
}

// PasteModel contains a number of attributes describing the paste in detail.
type PasteModel struct {
	Source     string    `json:"Source"`
	ID         string    `json:"Id"`
	Title      string    `json:"Title"`
	Date       time.Time `json:"Date"`
	EmailCount int       `json:"EmailCount"`
}

// Query contains a parameter which could be passed to the RESTful endpoint.
type Query struct {
	Parameter string
	Value     string
}
