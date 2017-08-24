# hibpgo

[![GoDoc](https://godoc.org/github.com/phlipse/hibpgo?status.svg)](https://godoc.org/github.com/phlipse/hibpgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/phlipse/hibpgo)](https://goreportcard.com/report/github.com/phlipse/hibpgo)

Package **hibpgo** provides access to the [Have I been Pwned?](https://haveibeenpwned.com/) API from Troy Hunt.
It supports all the RESTful service endpoints and parameters of API version 2.

## Usage

### Case-sensitivity
Accounts, names of breaches and domains are case-insensitive. Only passwords are looked up exactly as specified.

### User-Agent
The [Have I been Pwned?](https://haveibeenpwned.com/) API requires an user-agent set in each request header. The used user-agent can be retrieved through the variable *hibp.UserAgent*. It defaults to the string "hibpgo-<GO_VERSION>", for example "hibpgo-go/1.8.3".

### API rate limit
When performaing multiple requests sequentially (e.g. in a loop), *hibp.APIRateLimit* (type time.Duration) should be used as sleep time between each request. Otherwise the rate limit will be hit and punished by two seconds of sleep time by the API. This will be handled by the package itself, but the resulting penalty is higher then waiting the few milliseconds specified in *hibp.APIRateLimit*.

### Concurrency
This package can't be used by concurrent go routines.

## Documentation
Use [godoc](https://godoc.org/github.com/phlipse/hibpgo) for further documentation and more examples.

## Contributing
Pull requests are welcome!

## License

[MIT License](https://github.com/phlipse/hibpgo/blob/master/LICENSE)