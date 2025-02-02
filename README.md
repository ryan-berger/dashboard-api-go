# Meraki Dashboard API Go Library

The Meraki Dashboard API Golang library provides all current Meraki [dashboard API](https://developer.cisco.com/meraki/api-v1/) 
calls to interface with the Cisco Meraki cloud-managed platform.

Meraki generates the library based on dashboard API's OpenAPI spec to keep it up to date with the latest API releases.
The library requires Go 1.18+, receives support from the community, and you can install it according to the instructions below:

```shell
go get github.com/meraki/dashboard-api-go/client
```

## Features

While you can make direct HTTP requests to dashboard API in any programming language or REST API client, using a client library can make it easier for you to focus on your specific use case, without the overhead of having to write functions to handle the dashboard API calls.

* Support for all API endpoints, as it uses the [OpenAPI Generator](https://openapi-generator.tech) to generate source code from the Meraki [OpenAPI specification](https://api.meraki.com/api/v1/openapiSpec)
* Automatic retries upon 429 rate limit errors, using the [`Retry-After` field](https://developer.cisco.com/meraki/api-v1/#!rate-limit) within response headers


## Setup

1. Enable API access in your Meraki dashboard organization and obtain an API key ([instructions](https://documentation.meraki.com/zGeneral_Administration/Other_Topics/The_Cisco_Meraki_Dashboard_API))

2. Keep your API key safe and secure, as it is similar to a password for your dashboard. If publishing your Go code to a wider audience, please research secure handling of API keys.

3. Install the latest version of the [Go programming language](https://golang.org/doc/install)

4. Use `go get` command to install the library:
    * `go get github.com/meraki/dashboard-api-go/client`
   
5. Export your API key as an [environment variable](https://www.twilio.com/blog/2017/01/how-to-set-environment-variables.html), for example:

    ```shell
    export MERAKI_DASHBOARD_API_KEY=093b24e85df15a3e66f1fc359f4c48493eaa1b73
    ```

## Usage

Client API documentation is available in the **/client/docs** directory.
You can find fully working example scripts in the **/examples** folder.

### GetOrganizations Example

```go
package main

import "github.com/meraki/dashboard-api-go/client"
import "context"
import "fmt"
import "os"

func main() {
	configuration := client.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))

	apiClient := client.NewAPIClient(configuration)

	orgs, _, err := apiClient.ConfigureApi.GetOrganizations(context.Background()).Execute()

	if err != nil {
		fmt.Println("Meraki API call failed. Details: ", err)
	}

	for _, org := range orgs {
		fmt.Printf("%s\n", *org.Name)
	}
}
```


## Note for application developers and ecosystem partners

We're so glad that you're leveraging our Go library. It's best practice to identify your application with every API request that you make. 
You can easily do this automatically by setting:

```go
configuration.UserAgent = "MY-APPLICATION-IDENTIFIER"
```

Unless you are an ecosystem partner, this identifier is optional.

1. If you are an ecosystem partner and you have questions about this requirement, please reach out to your ecosystem rep.
2. If you have any questions about the formatting, please ask your question by opening an issue in this repo.