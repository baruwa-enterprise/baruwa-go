# baruwa-go

## Golang bindings for the Baruwa REST API

[![Build Status](https://travis-ci.org/baruwa-enterprise/baruwa-go.svg?branch=master)](https://travis-ci.org/baruwa-enterprise/baruwa-go)
[![codecov](https://codecov.io/gh/baruwa-enterprise/baruwa-go/branch/master/graph/badge.svg)](https://codecov.io/gh/baruwa-enterprise/baruwa-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/baruwa-enterprise/baruwa-go)](https://goreportcard.com/report/github.com/baruwa-enterprise/baruwa-go)
[![GoDoc](https://godoc.org/github.com/baruwa-enterprise/baruwa-go?status.svg)](https://godoc.org/github.com/baruwa-enterprise/baruwa-go)
[![MPLv2 License](https://img.shields.io/badge/license-MPLv2-blue.svg?style=flat-square)](https://www.mozilla.org/MPL/2.0/)

```
Usage: baruwa -k -s COMMAND [arg...]

A cmdline client for the Baruwa REST API.
                     
Options:             
  -k, --api-token    Baruwa API OAUTH Token (env $BARUWA_API_TOKEN)
  -s, --server-url   Baruwa server url (env $BARUWA_API_SERVER)
                     
Commands:            
  user               manage user accounts
  users              list user accounts
  domain             manage domains
  domains            list domains
  organization       manage organizations
  organizations      list organizations
  systemstatus       show system status
                     
Run 'baruwa COMMAND --help' for more information on a command.
```

## Requirements

* Golang 1.10.x or higher

## Installation

```console
$ go get github.com/baruwa-enterprise/baruwa-go
```

## Testing

``make test``

## Contributing

1. Fork it (https://github.com/baruwa-enterprise/baruwa-go/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## License

All code is licensed under the
[MPLv2 License](https://github.com/baruwa-enterprise/baruwa-go/blob/master/LICENSE).
