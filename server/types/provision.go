package types

type PipelineProvision struct {
	PipelineId   string               `json:"pipelineId"`
	Capabilities PipelineCapabilities `json:"capabilities"`
}
