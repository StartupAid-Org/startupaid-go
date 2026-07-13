# startupaid-go

The official Go client for the [startupAid](https://startupaid.org) API — send
transactional email, send push, schedule social posts, and convert currency, all
with your account API key.

## Install

```bash
go get github.com/StartupAid-Org/startupaid-go
```

```go
import startupaid "github.com/StartupAid-Org/startupaid-go"
```

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	startupaid "github.com/StartupAid-Org/startupaid-go"
)

func main() {
	client := startupaid.New("sk_your_key")
	ctx := context.Background()

	res, err := client.Convert(ctx, "USD", "NGN", 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("100 USD = %.2f NGN\n", res.Result)
}
```

Create an API key in your startupAid dashboard → **API Keys**.

## Usage

### Email

```go
msg, err := client.SendEmail(ctx, startupaid.SendEmailRequest{
	From:    "Acme <hi@acme.com>",
	To:      "user@example.com",
	Subject: "Welcome",
	Body:    "<h1>Hi 👋</h1>",
})

status, err := client.GetMessage(ctx, msg.ID)
```

### Currency

```go
res, err := client.Convert(ctx, "USD", "EUR", 250)
rates, err := client.Rates(ctx, "USD")
list, err := client.Currencies(ctx)
```

### Push

```go
_, err := client.SendPush(ctx, "app_id", startupaid.SendPushRequest{
	Target: startupaid.PushTarget{UserRef: "u_123"},
	Title:  "Your order shipped",
	Body:   "Track it in the app.",
})
```

### Social

```go
channels, err := client.ListChannels(ctx)

post, err := client.SchedulePost(ctx, startupaid.SchedulePostRequest{
	Channels: []string{channels[0].ID},
	Content:  "We just shipped 🚀",
})

posts, err := client.ListPosts(ctx)
```

### Audiences

```go
lists, err := client.ListAudiences(ctx)
err = client.AddContact(ctx, "audience_id", startupaid.AddContactRequest{
	Email:     "new@example.com",
	FirstName: "Ada",
})
```

## Configuration

```go
client := startupaid.New("sk_your_key",
	startupaid.WithBaseURL("http://localhost:8080"), // self-host / testing
	startupaid.WithHTTPClient(customHTTPClient),      // custom timeouts/transport
)
```

## Errors

Non-2xx responses return an `*APIError`:

```go
_, err := client.Convert(ctx, "USD", "NGN", 1)
var apiErr *startupaid.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Message)
}
```

Each method requires your account to be subscribed to the matching product
(email, currency, push, social); otherwise the API returns an upgrade prompt.

## License

MIT
