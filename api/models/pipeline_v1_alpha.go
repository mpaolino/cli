package models

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

type PipelineV1Alpha struct {
	Pipeline struct {
		ID     string `json:"ppl_id"`
		Name   string `json:"name,omitempty"`
		State  string `json:"state,omitempty"`
		Result string `json:"result,omitempty" yaml:"result,omitempty"`
		Reason string `json:"result_reason,omitempty" yaml:"result_reason,omitempty"`
		Error  string `json:"error_description,omitempty"  yaml:"error_description,omitempty"`
	} `json:"pipeline,omitempty"`
}

type PipelineV1AlphaPartialRebuildRequest struct {
	ScheduleType string `json:"schedule_type"`
	PplToRebuild string `json:"ppl_to_rebuild"`
	RequestToken string `json:"request_token"`
}

func NewPipelineV1AlphaFromJson(data []byte) (*PipelineV1Alpha, error) {
	j := PipelineV1Alpha{}

	err := json.Unmarshal(data, &j)

	if err != nil {
		return nil, err
	}

	return &j, nil
}

func (j *PipelineV1Alpha) ToYaml() ([]byte, error) {
	return yaml.Marshal(j)
}

func (j *PipelineV1Alpha) IsDone() bool {
	return j.Pipeline.State == "done"
}

func NewPipelineV1AlphaPartialRebuildRequest(scheduleType string, pplId string, reqToken string) (*PipelineV1AlphaPartialRebuildRequest, error) {
	j := PipelineV1AlphaPartialRebuildRequest{}
	j.ScheduleType = scheduleType
	j.PplToRebuild = pplId
	j.RequestToken = reqToken
	return &j, nil
}
