package models

import "encoding/json"

type WorkflowListV1Alpha struct {
	Workflow []WorkflowV1Alpha `json:"workflows" yaml:"projects"`
}

func NewWorkflowListV1AlphaFromJson(data []byte) (*WorkflowListV1Alpha, error) {
	list := []WorkflowV1Alpha{}

	err := json.Unmarshal(data, &list)

	if err != nil {
		return nil, err
	}

	for _, p := range list {
		if p.ApiVersion == "" {
			p.ApiVersion = "v1alpha"
		}

		if p.Kind == "" {
			p.Kind = "Project"
		}
	}

	return &WorkflowListV1Alpha{Workflow: list}, nil
}
