<div id="top"></div>

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Go Reference][reference-shield]][reference-url]
[![Go Report Card][go-report-card-shield]][go-report-card-url]
[![Apache 2.0 License][license-shield]][license-url]
[![Workflow][workflow-shield]][workflow-url]

<br />
<div align="center">
  <a href="https://github.com/smartpay-co/sdk-go">
		<picture>
			<source media="(prefers-color-scheme: dark)" srcset="https://assets.smartpay.co/logo/banner/smartpay-logo-dark.png" />
			<source media="(prefers-color-scheme: light)" srcset="https://assets.smartpay.co/logo/banner/smartpay-logo.png" />
			<img alt="Smartpay" src="https://assets.smartpay.co/logo/banner/smartpay-logo.png" style="width: 797px;" />
		</picture>
  </a>

  <p align="center">
    <a href="https://docs.smartpay.co/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/smartpay-co/integration-examples">View Example</a>
    ·
    <a href="https://github.com/smartpay-co/sdk-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/smartpay-co/sdk-go/issues">Request Feature</a>
  </p>
</div>

# Smartpay Go SDK

Smartpay Go SDK offers easy access to Smartpay API from applications written in Go.

## Prerequisites

Go v1.18+

## Installation

```shell
go get github.com/smartpay-co/sdk-go
```

## Usage

Let's give a quick example first about how to create a new checkout session.

```golang
package main

import (
	"context"
	"fmt"
	. "github.com/smartpay-co/sdk-go"
)

func main() {
	ctx := context.Background()
	client, _ := NewClientWithResponses("<YOUR_SECRET_KEY>", "<YOUR_PUBLIC_KEY>")

	checkoutPayload := CreateACheckoutSessionJSONRequestBody{
		Currency: CurrencyJPY,
		Amount:   350,
		Items: []Item{
			{Name: "レブロン 18 LOW", Amount: 1000, Currency: CurrencyJPY, Quantity: Ptr(1)},
			{Name: "discount", Amount: 100, Currency: CurrencyJPY, Kind: Ptr(LineItemKindDiscount)},
			{Name: "tax", Amount: 250, Currency: CurrencyJPY, Kind: Ptr(LineItemKindTax)},
		},
		ShippingInfo: ShippingInfo{
			Address:     Address{Country: AddressCountryJP, PostalCode: "123", Locality: "locality", Line1: "line1"},
			FeeAmount:   Ptr(float32(100)),
			FeeCurrency: Ptr(CurrencyJPY),
		},
		CaptureMethod: Ptr(CaptureMethodManual),
		Reference:     Ptr("order_ref_1234567"),
		SuccessUrl:    "https://smartpay.co",
		CancelUrl:     "https://smartpay.co",
	}

	result, err := client.CreateACheckoutSessionWithResponse(ctx, checkoutPayload)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result.Body))
}
```

### Initialize A New Client

You need to get your API key to initialize a new client. You can find it on your [dashboard](https://dashboard.smartpay.co/settings/credentials).

You can also customize the client by passing function options into the constructors.

```golang
client, _ := NewClientWithResponses("<SECRET_KEY>", WithHTTPClient(&httpClient), WithBaseURL("http://localhost:9000"))
```

### Invoke an API Point

Our clients supprts struct parameters and responses for API calls. In most cases you'll be satisfied with this method. Take the example from above, `CreateACheckoutSessionWithResponse` is so much easier to use with the predefined request and response objects.

But if you want full control for maximium flexibility, then you can use `CreateACheckoutSessionWithBody` and handle the request and response on your own.

This design applies to all our APIs.

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/smartpay-co/sdk-go.svg
[contributors-url]: https://github.com/smartpay-co/sdk-go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/smartpay-co/sdk-go.svg
[forks-url]: https://github.com/smartpay-co/sdk-go/network/members
[stars-shield]: https://img.shields.io/github/stars/smartpay-co/sdk-go.svg
[stars-url]: https://github.com/smartpay-co/sdk-go/stargazers
[issues-shield]: https://img.shields.io/github/issues/smartpay-co/sdk-go.svg
[issues-url]: https://github.com/smartpay-co/sdk-go/issues
[license-shield]: https://img.shields.io/github/license/smartpay-co/sdk-go.svg
[license-url]: https://github.com/smartpay-co/sdk-go/blob/main/LICENSE.txt
[reference-shield]: https://pkg.go.dev/badge/github.com/smartpay-co/sdk-go.svg
[reference-url]: https://pkg.go.dev/github.com/smartpay-co/sdk-go
[go-report-card-shield]: https://goreportcard.com/badge/github.com/smartpay-co/sdk-go
[go-report-card-url]: https://goreportcard.com/report/github.com/smartpay-co/sdk-go
[workflow-shield]: https://github.com/smartpay-co/sdk-go/actions/workflows/ci.yml/badge.svg
[workflow-url]: https://github.com/smartpay-co/sdk-go/actions/workflows/ci.yml
