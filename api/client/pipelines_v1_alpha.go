package client

import (
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	models "github.com/semaphoreci/cli/api/models"
)

type PipelinesApiV1AlphaApi struct {
	BaseClient           BaseClient
	ResourceNameSingular string
	ResourceNamePlural   string
}

func NewPipelinesV1AlphaApi() PipelinesApiV1AlphaApi {
	baseClient := NewBaseClientFromConfig()
	baseClient.SetApiVersion("v1alpha")

	return PipelinesApiV1AlphaApi{
		BaseClient:           baseClient,
		ResourceNamePlural:   "pipelines",
		ResourceNameSingular: "pipeline",
	}
}

func (c *PipelinesApiV1AlphaApi) DescribePpl(id string) (*models.PipelineV1Alpha, error) {
	body, status, err := c.BaseClient.Get(c.ResourceNamePlural, id)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("connecting to Semaphore failed '%s'", err))
	}

	if status != 200 {
		return nil, errors.New(fmt.Sprintf("http status %d with message \"%s\" received from upstream", status, body))
	}

	return models.NewPipelineV1AlphaFromJson(body)
}

func (c *PipelinesApiV1AlphaApi) StopPpl(id string) ([]byte, error) {
	request_body := []byte("{\"terminate_request\": true}")

	body, status, err := c.BaseClient.Patch(c.ResourceNamePlural, id, request_body)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("connecting to Semaphore failed '%s'", err))
	}

	if status != 200 {
		return nil, errors.New(fmt.Sprintf("http status %d with message \"%s\" received from upstream", status, body))
	}

	return body, nil
}

func (c *PipelinesApiV1AlphaApi) CreatePpl(id string) ([]byte, error) {
	requestToken, err := uuid.NewV1()

	req, err := models.NewPipelineV1AlphaPartialRebuildRequest("partial_rebuild", id, requestToken.String())

	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating reuild request failed '%s'", err))
	}

	d, _ := json.Marshal(req)
	fmt.Printf("aaaaaaaaa %s\n", string(d))
	requestBody := fmt.Sprintf("%s", string(d))
	body, status, err := c.BaseClient.Post(c.ResourceNamePlural, []byte(requestBody))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("connecting to Semaphore failed '%s'", err))
	}

	if status != 200 {
		return nil, errors.New(fmt.Sprintf("http status %d with message \"%s\" received from upstream", status, body))
	}

	return body, nil
}
