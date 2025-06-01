package notification

import (
	"log"
	"os"
)

type NotificationProvider interface {
	Notify(message string) error
}

// Providers are injected in this function
func InitProviders() []NotificationProvider {
	return []NotificationProvider{
		Discord{WebhookURL: os.Getenv("DISCORD_WEBHOOK")},
		Dummy{},
	}
}

func Run(message string, providers []NotificationProvider) []error {
	errors := []error{}

	for _, provider := range providers {
		err := provider.Notify(message) // TODO: Should be async
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func ProcessErrors(errors []error) {
	for _, err := range errors {
		log.Printf("ERROR: %v\n", err)
	}
}
