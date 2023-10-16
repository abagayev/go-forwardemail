# Go client for Forward Email

[![Go Reference](https://pkg.go.dev/badge/github.com/namecheap/go-namecheap-sdk.svg)](https://pkg.go.dev/github.com/namecheap/go-namecheap-sdk/v2)
[![CodeQL](https://github.com/MichaelCurrin/badge-generator/workflows/CodeQL/badge.svg)](https://github.com/abagayev/go-forwardemail/actions?query=workflow%3ACodeQL "Code quality workflow status")
[![codecov](https://codecov.io/gh/abagayev/go-forwardemail/graph/badge.svg?token=5JZDkzKaGf)](https://codecov.io/gh/abagayev/go-forwardemail)

- [Forward Email API Documentation](https://forwardemail.net/en/email-api)
- [Forward Email Terraform Provider](https://github.com/abagayev/terraform-provider-forwardemail)

### How to install

```shell
$ go get github.com/abagayev/go-forwardemail
```

### Basic usage

```go
import "github.com/abagayev/go-forwardemail/forwardemail"

client := forwardemail.NewClient(forwardemail.ClientOptions{
ApiKey: key,
})

account, err := client.GetAccount()
```

### Contribution

Feel free to add comments, issues, pull requests or buy me a coffee:  
https://www.buymeacoffee.com/tonybug
