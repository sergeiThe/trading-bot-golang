package notification

import (
	"log"
)

type Dummy struct {}

func (n Dummy) Notify(message string) error {
	log.Println("Dummy notified")
	return nil
}
