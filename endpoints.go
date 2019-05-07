// Copyright (c) 2017 Philipp Weber
// Use of this source code is governed by the MIT license
// which can be found in the repositorys LICENSE file.

package hibpgo

import (
	"encoding/json"
)

// BreachedAccount returns all breaches which affected the account.
func BreachedAccount(account, domain string, truncateResponse, includeUnverified bool) ([]BreachModel, error) {
	querys := []Query{}
	if domain != "" {
		querys = append(querys, Query{Parameter: "domain", Value: domain})
	}
	if truncateResponse {
		querys = append(querys, Query{Parameter: "truncateResponse", Value: "true"})
	}
	if includeUnverified {
		querys = append(querys, Query{Parameter: "includeUnverified", Value: "true"})
	}

	content, err := callEndpoint("breachedaccount", account, querys)
	if err != nil {
		return nil, err
	}

	var b []BreachModel
	err = json.Unmarshal(content, &b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Breaches returns all known breaches, may be filtered by domain name.
func Breaches(domain string) ([]BreachModel, error) {
	querys := []Query{}
	if domain != "" {
		querys = append(querys, Query{Parameter: "domain", Value: domain})
	}

	content, err := callEndpoint("breaches", "", querys)
	if err != nil {
		return nil, err
	}

	var b []BreachModel
	err = json.Unmarshal(content, &b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Breach returns specific breach.
func Breach(name string) (BreachModel, error) {
	content, err := callEndpoint("breach", name, []Query{})
	if err != nil {
		return BreachModel{}, err
	}

	var b BreachModel
	err = json.Unmarshal(content, &b)
	if err != nil {
		return BreachModel{}, err
	}

	return b, nil
}

// Dataclasses returns all known dataclasses.
func Dataclasses() ([]string, error) {
	content, err := callEndpoint("dataclasses", "", []Query{})
	if err != nil {
		return nil, err
	}

	var s []string
	err = json.Unmarshal(content, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// PastedAccount returns all pastes which affected the account.
func PastedAccount(account string) ([]PasteModel, error) {
	content, err := callEndpoint("pasteaccount", account, []Query{})
	if err != nil {
		return nil, err
	}

	var p []PasteModel
	err = json.Unmarshal(content, &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// PwnedPassword checks if password is compromised due to a data breach.
func PwnedPassword(password string, originalPasswordIsAHash bool) (bool, error) {
	querys := []Query{}
	if originalPasswordIsAHash {
		querys = append(querys, Query{Parameter: "originalPasswordIsAHash", Value: "true"})
	}

	content, err := callEndpoint("pwnedpassword", password, querys)

	if err != nil {
		// an error occurred
		return false, err
	} else if string(content) == password {
		// password was found
		return true, nil
	} else {
		// everything ok
		return false, nil
	}
}
