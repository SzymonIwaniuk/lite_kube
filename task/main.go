package task

import (
	"github.com/docker/docker/client"
)

func main() {
	c := Config{
		Name:  "test-container-1",
		Image: "postgres:18",
	}

	dc, _ := client.NewClientWithOpts(client.FromEnv)
	d := Docker{
		Client: dc,
		Config: c,
	}

	d.Run()
}
