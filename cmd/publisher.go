/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gocloud.dev/pubsub"
)

// publisherCmd represents the publisher command
var publisherCmd = &cobra.Command{
	Use:   "publisher",
	Short: "message publisher",
	Long:  `publish messages to a given topic, can be from any cloud provider`,
	Run:   runPublisher,
}

func init() {
	publisherCmd.PersistentFlags().StringP("message", "m", "hello World", "Message to publish")
	viper.BindPFlag("message", publisherCmd.PersistentFlags().Lookup("message"))

	publisherCmd.PersistentFlags().StringP("url", "u", "Topic URL", "topic to publish to")
	viper.BindPFlag("url", publisherCmd.PersistentFlags().Lookup("url"))

	rootCmd.AddCommand(publisherCmd)
}

func runPublisher(cmd *cobra.Command, args []string) {
	if !viper.IsSet("url") {
		fmt.Println("Please set url")
		return
	}

	url := viper.GetString("url")

	// Open the topic
	topic, err := pubsub.OpenTopic(context.Background(), url)
	if err != nil {
		log.Fatalf("Failed to open topic: %v", err)
	}
	defer topic.Shutdown(context.Background())

	// Publish a message
	err = topic.Send(context.Background(), &pubsub.Message{
		Body: []byte(viper.GetString("message")),
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message published!")
}
