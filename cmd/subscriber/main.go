package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/pubsub"
)

var (
	//projectID of GCP project.
	projectID = "my-pubsub-404318" // TODO: get from .json
	// subID of PubSub subscription.
	subID = "MyTopic-sub"
)

// publishHandler handler for subscribe message.
func handler(ctx context.Context) error {
	// Init PubSub Client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Print("GOOGLE_APPLICATION_CREDENTIALS:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		return err
	}
	defer client.Close()

	// Create a subscription references to a topic.
	sub := client.Subscription(subID)

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		log.Printf("got data: %q\n", string(msg.Data))
		log.Printf("got attribute: %q\n", msg.Attributes)
		msg.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/workspace/go-pubsub-ws/.service_account.json") // TODO: Why did I need to do this?

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start running the server.
	go func() {
		err := handler(ctx)
		// if err != nil && errors.Is(err, context.Canceled) {
		if err != nil {
			log.Fatalf("failed to pull messages: %v\n", err)
		}
	}()
	log.Println("subscriber is starting...")

	// Receive signal to shutdown the server.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Printf("signal %d received, shutting down gracefully...", <-quit)
	cancel()

	log.Println("finished graceful shutdown")
}
