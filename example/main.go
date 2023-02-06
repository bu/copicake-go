package main

import (
	"log"
	"os"
	"time"

	"github.com/bu/copicake-go/v1"
)

func main() {
	// create a new copicake client
	cake, err := copicake.New(&copicake.ClientConfig{
		ApiKey:        "YOUR_API_KEY",
		RetryTimeout:  1 * time.Second,
		RetryMaxTries: 5,
	})

	if err != nil {
		log.Fatalf("cannot init Copicake client: %s", err)
	}

	// create new render job
	job, err := cake.NewRenderRequest(copicake.RenderRequest{
		TemplateID: "YOUR_TEMPLATE_ID", // Get template ID
		Changes: []copicake.C{
			{
				"name": "text-message", // change to your text layer name
				"text": "Test",         // change to your text
			},
		},
	})

	if err != nil {
		log.Fatalf("cannot create Copicake render job: %s", err)
	}

	// query current process status
	status, err := job.Status()

	if err != nil {
		log.Fatalf("cannot get process status")
	}

	log.Printf("status: %+v", status)

	// wait for image to be ready, return result url
	url, err := job.URL()

	if err != nil {
		log.Fatalf("cannot get image URL: %s", err)
	}

	log.Printf("url: %+v", url)

	// wait for image to be ready, return result image content
	image, err := job.Image()

	if err != nil {
		log.Fatalf("cannot get image: %s", err)
	}

	// dont output to stdout, save to file
	err = os.WriteFile("test.png", image, 0644)
	if err != nil {
		log.Fatalf("cannot save image: %s", err)
	}
}
