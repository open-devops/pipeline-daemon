package types

type PipelineProvision struct {
	PipelineId   string               `json:"pipelineId"`
	PipelineName string               `json:"pipelineName"`
	Capabilities PipelineCapabilities `json:"capabilities"`
}
