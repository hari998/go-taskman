package notify

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
)

func Show(title, message string) {

	if os.Getenv("WSL_DISTRO_NAME") != "" {
		return
	}

	err := beeep.Notify(title, message, "icon.png")

	if err != nil {
		fmt.Println("Error showing notification:", err)
	}
}
