package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func Publish(projectID, topicID, msg string) error {
	// projectID := "my-project-id"
	// topicID := "my-topic"
	// msg := "Hello World"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

var (
	topic = flag.String("topic", "", "Set topic")
	proj  = flag.String("project", "", "set project")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("project: %s | topic: %s\n", *proj, *topic)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Msg: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			err := Publish(*proj, *topic, text)
			if err != nil {
				log.Fatalf("err: %v", err)
			}
			fmt.Println("published msg: ", text)
		} else {
			// exit if user entered an empty string
			break
		}

	}
	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

}
