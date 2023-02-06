# copicake-go
🍰 Copicake, a data-driven image generating service to let you generate any social media material with just ONE CLICK.

* 🔗 Website: https://copicake.com/
* 📘 Official API Docs: https://docs.copicake.com/

# Installations
Run `go get github.com/bu/copicake-go` under your module directory

# Usage
Detailed checkout example codes under `example/`, here is a quick example:

```go
package main

import (
	"log"
	"time"

	"github.com/bu/copicake-go/v1"
)

func main() {
	// create a new copicake client
	cake, err := copicake.New(&copicake.ClientConfig{
		ApiKey:        "<YOUR_API_KEY_HERE>",
		RetryTimeout:  1 * time.Second,
		RetryMaxTries: 5,
	})

	if err != nil {
		log.Fatalf("cannot init Copicake client: %s", err)
	}

	// create new render job
	job, err := cake.NewRenderRequest(copicake.RenderRequest{
		TemplateID: "", // Get template ID
		Changes: []copicake.C{
			{
				"name": "message",
				"text": "2",
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

	log.Printf("image: %+v", image)
}
```