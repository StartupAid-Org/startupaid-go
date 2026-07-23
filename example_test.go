package startupaid_test

import (
	"context"
	"fmt"
	"log"
	"time"

	startupaid "github.com/StartupAid-Org/startupaid-go"
)

func Example() {
	client := startupaid.New("sk_your_key")

	res, err := client.Convert(context.Background(), "USD", "NGN", 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("100 USD = %.2f NGN\n", res.Result)
}

func ExampleClient_SendEmail() {
	client := startupaid.New("sk_your_key")

	msg, err := client.SendEmail(context.Background(), startupaid.SendEmailRequest{
		From:    "Acme <hi@acme.com>",
		To:      "user@example.com",
		Subject: "Welcome aboard",
		Body:    "<h1>Hi there 👋</h1>",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sent:", msg.ID)
}

func ExampleClient_SendOTP() {
	client := startupaid.New("sk_your_key")
	ctx := context.Background()

	ch, err := client.SendOTP(ctx, startupaid.SendOTPRequest{
		To:      "+2348012345678",
		AppName: "Acme",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sent:", ch.ID)

	ok, err := client.VerifyOTP(ctx, "+2348012345678", "123456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verified:", ok)
}

func ExampleClient_SchedulePost() {
	client := startupaid.New("sk_your_key")
	ctx := context.Background()

	channels, err := client.ListChannels(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var ids []string
	for _, ch := range channels {
		ids = append(ids, ch.ID)
	}

	at := time.Now().Add(2 * time.Hour)
	post, err := client.SchedulePost(ctx, startupaid.SchedulePostRequest{
		Channels:     ids,
		Content:      "We just shipped multi-region sending 🚀",
		ScheduledFor: &at,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("scheduled:", post.ID)
}
