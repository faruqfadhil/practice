// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	topic        = flag.String("topic", "", "Set topic")
	proj         = flag.String("project", "", "set project")
	action       = flag.String("action", "create-topic", "set action")
	subscription = flag.String("subscription", "", "set subs")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("action: %s | project: %s | topic: %s\n", *action, *proj, *topic)
	ctx := context.Background()
	// Creates a client.
	client, err := pubsub.NewClient(ctx, *proj)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	switch *action {
	case "create-topic":
		// Creates the new topic.
		topic, err := client.CreateTopic(ctx, *topic)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}

		fmt.Printf("Topic %v created.\n", topic)
	case "create-subscription":
		t := client.Topic(*topic)
		sub, err := client.CreateSubscription(ctx, *subscription, pubsub.SubscriptionConfig{
			Topic:       t,
			AckDeadline: 20 * time.Second,
		})
		if err != nil {
			log.Fatalf("CreateSubscription: %v", err)
		}
		fmt.Printf("Created subscription: %v\n", sub)
	}
}
