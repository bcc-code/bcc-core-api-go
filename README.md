# bcc-core-api-go
The Go SDK for the BCC Core API

-------------------------------------

## Documentation

- [Godoc](https://pkg.go.dev/github.com/bcc-code/bcc-core-api-go/coreapi).

## Getting started

### Installation

```sh
go get github.com/bcc-code/bcc-core-api-go
```

### Usage

Create an Core API client by providing the details of your client.

```go
import (
	"context"
	"fmt"
	"log"

	"github.com/bcc-code/bcc-core-api-go/coreapi"
	"github.com/google/uuid"
)

func main() {
	clientID := "EXAMPLE_16L9d34h0qe4NVE6SaHxZEid"
	clientSecret := "EXAMPLE_XSQGmnt8JdXs23407hrK6XXXXXXX"

	// Initialize a new client using a domain, client ID and client secret.
	coreAPI, err := coreapi.New(
		coreapi.UrlSandbox, // Predefined Core API url for the Sandbox environment
		coreapi.WithClientCredentials(
			context.Background(),
			coreapi.CredEnvSandbox, // Predefined Identity server information for the Sandbox environment
			clientID,
			clientSecret,
		),
	)
	if err != nil {
		log.Fatalf("failed to initialize the core API client: %+v", err)
	}

	// Now we can interact with the BCC Core API.
	// Fetch a person by Uid

	p, err := coreAPI.Person.Get(context.Background(),
		uuid.MustParse("657a66ca-9cd0-4b61-9476-697016e26fbc"))

	if err != nil {
		log.Fatalf("failed to get a person by Uid: %+v", err)
	}

	fmt.Println(p.DisplayName)
}

```
