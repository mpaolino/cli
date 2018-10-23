package models

type WorkflowV1Alpha struct {
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion"`
	Kind       string `json:"kind,omitempty" yaml:"kind"`
	Metadata   struct {
		Name string `json:"name,omitempty"`
		Id   string `json:"id,omitempty"`
	} `json:"metadata,omitempty"`

	Spec struct {
		Repository struct {
			Url string `json:"url,omitempty"`
		} `json:"repository,omitempty"`
	} `json:"spec,omitempty"`
}
