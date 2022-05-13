package main
import (
	"log"

	"github.com/Rinshad/webhook-repo/webhook/api"
)

func main() {
	err := api.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
