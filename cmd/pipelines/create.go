package pipelines

import (
	"fmt"

	client "github.com/semaphoreci/cli/api/client"
	"github.com/semaphoreci/cli/cmd/utils"
)

func Create(id string) {
	c := client.NewPipelinesV1AlphaApi()
	body, err := c.CreatePpl(id)
	utils.Check(err)
	fmt.Printf("%s\n", string(body))
}
