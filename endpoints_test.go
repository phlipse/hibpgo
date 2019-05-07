// Copyright (c) 2019 Philipp Weber. All rights reserved.
// Use of this source code is governed by the MIT license
// which can be found in the repositorys LICENSE file.

package hibpgo

import "testing"

func TestBreachedAccount(t *testing.T) {
	b, err := BreachedAccount("test@example.com")
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if len(b) == 0 {
		t.Errorf("expected breaches got none")
	}
}

func TestBreachedAccountOpt(t *testing.T) {
	b, err := BreachedAccountOpt("test@example.com", "", false, false)
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if len(b) == 0 {
		t.Errorf("expected breaches got none")
	}
}

func TestBreaches(t *testing.T) {
	b, err := Breaches()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if len(b) == 0 {
		t.Errorf("expected breaches got none")
	}
}

func TestBreachesOpt(t *testing.T) {
	b, err := BreachesOpt("adobe.com")
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if len(b) == 0 {
		t.Errorf("expected breaches got none")
	}
}

func TestBreach(t *testing.T) {
	b, err := Breach("Adobe")
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if b.Name == "" {
		t.Errorf("expected breach got none")
	}
}

func TestDataclasses(t *testing.T) {
	d, err := Dataclasses()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if len(d) == 0 {
		t.Errorf("expected dataclasses got none")
	}
}

func TestPastedAccount(t *testing.T) {
	p, err := PastedAccount("test@example.com")
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if len(p) == 0 {
		t.Errorf("expected pastes got none")
	}
}

func TestPwnedPassword(t *testing.T) {
	p, err := PwnedPassword("12 34")
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if !p {
		t.Errorf("expected true got %t", p)
	}
}

func TestPwnedPasswordOpt(t *testing.T) {
	p, err := PwnedPasswordOpt("12 34", false)
	if err != nil {
		t.Errorf("got error: %s", err.Error())
	}
	if !p {
		t.Errorf("expected true got %t", p)
	}
}
