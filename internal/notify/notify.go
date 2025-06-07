package notify

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func Show(title, message string) {
	err := beeep.Notify(title, message, "icon.png")

	if err != nil {
		fmt.Println("Error showing notification:", err)
	}
}