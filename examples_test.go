// Copyright 2017 (c) Philipp Weber. All rights reserved.
// Use of this source code is governed by the MIT license
// which can be found in the repositorys LICENSE file.

package hibpgo_test

import (
	"fmt"
	"time"

	"github.com/phlipse/hibpgo"
)

func ExampleBreachedAccount() {
	b, err := hibpgo.BreachedAccount("test@example.com", "000webhost.com", false, false)
	fmt.Println(b[0].Name, err)
	// Output:
	// 000webhost <nil>
}

func ExampleBreaches() {
	b, err := hibpgo.Breaches("adobe.com")
	fmt.Println(b[0].Domain, err)
	// Output:
	// adobe.com <nil>
}

func ExampleBreach() {
	b, err := hibpgo.Breach("Adobe")
	fmt.Println(b.Name, err)
	// Output:
	// Adobe <nil>
}

func ExampleDataclasses() {
	d, err := hibpgo.Dataclasses()
	fmt.Println(d[0], err)
	// Output:
	// Account balances <nil>
}

func ExamplePasteAccount() {
	p, err := hibpgo.PasteAccount("test@example.com")
	fmt.Println(p[0].Source, err)
	// Output:
	// Pastebin <nil>
}

func ExamplePwnedPassword() {
	// look up password
	p, err := hibpgo.PwnedPassword("Pa$$w0rd", false)
	fmt.Println(p, err)

	time.Sleep(hibpgo.APIRateLimit)

	// look up SHA1 hash of password
	p, err = hibpgo.PwnedPassword("353e8061f2befecb6818ba0c034c632fb0bcae1b", false)
	fmt.Println(p, err)

	time.Sleep(hibpgo.APIRateLimit)

	// look up password which itself looks like a SHA1 hash
	p, err = hibpgo.PwnedPassword("353e8061f2befecb6818ba0c034c632fb0bcae1b", true)
	fmt.Println(p, err)
	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
}
