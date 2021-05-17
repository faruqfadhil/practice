package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
)

func PullMsgs(projectID, subID string) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}

func init() {
	flag.Parse()
}

var (
	subs = flag.String("subs", "", "Set subscription")
	proj = flag.String("proj", "", "set project")
)

func main() {
	fmt.Printf("project: %s | subscription: %s\n", *proj, *subs)
	err := PullMsgs(*proj, *subs)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
