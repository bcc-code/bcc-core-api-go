# bcc-core-api-go
The Go SDK for the BCC Core API

-------------------------------------

## Documentation

- [Godoc](https://pkg.go.dev/github.com/bcc-code/bcc-core-api-go/bcccoreapi).

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

	"github.com/bcc-code/bcc-core-api-go/bcccoreapi"
)

func main() {
	clientID := "EXAMPLE_CLIENT_ID"
	clientSecret := "EXAMPLE_CLIENT_SECRET"

	// Initialize a new client using a domain, client ID and client secret.
	client, err := bcccoreapi.NewClient(
		bcccoreapi.WithClientCredentials(
			context.Background(),
			clientID,
			clientSecret,
		),
	)
	if err != nil {
		log.Fatalf("failed to initialize the core API client: %+v", err)
	}

	// Now we can interact with the BCC Core API.
	// Fetch a person by Uid

	p, err := client.Person.Get(context.Background(), "657a66ca-9cd0-4b61-9476-697016e26fbc")

	if err != nil {
		log.Fatalf("failed to get a person by Uid: %+v", err)
	}

	fmt.Println(p.DisplayName)
}

```
