/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/gcppubsub"
)

const (
	maxWorkers = 10
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "consumer",
	Long:  `pubsub consumer which reads from the given topic, supports many cloud provider`,
	Run:   runConsumer,
}

func init() {
	consumerCmd.PersistentFlags().StringP("topic", "t", "Topic URL", "topic to publish to")
	viper.BindPFlag("topic", consumerCmd.PersistentFlags().Lookup("topic"))

	rootCmd.AddCommand(consumerCmd)

}

func runConsumer(cmd *cobra.Command, args []string) {
	url := viper.GetString("topic")
	// Open the subscription
	sub, err := pubsub.OpenSubscription(context.Background(), url)
	if err != nil {
		log.Fatalf("Failed to open subscription: %v", err)
	}
	defer sub.Shutdown(context.Background())

	// Use a WaitGroup to track when all goroutines are done.
	var wg sync.WaitGroup

	workQueue := make(chan *pubsub.Message)

	// Start maxWorkers number of goroutines to process messages from the work queue
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range workQueue {
				processMessage(msg)
			}
		}()
	}

	// Create a channel to listen for the interrupt signal (CTRL+C)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Run the main loop in a goroutine so that it can be stopped by the interrupt signal
	go func() {
		for {
			msg, err := sub.Receive(context.Background())
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				return
			}
			workQueue <- msg
		}
	}()

	// Wait for the interrupt signal
	<-signals
	log.Println("Shutdown signal received, closing the subscription...")

	// Shut down the subscription and close the work queue
	if err := sub.Shutdown(context.Background()); err != nil {
		log.Printf("Failed to close subscription: %v", err)
	}
	close(workQueue)

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("All goroutines have finished, exiting...")
}

// processMessage processes a message.
func processMessage(msg *pubsub.Message) {
	fmt.Printf("Processing message: %s\n", msg.Body)
	msg.Ack()
}
