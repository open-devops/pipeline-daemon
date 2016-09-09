package types

type PipelineInfo struct {
	PipelineId          string       `json:"pipelineId"`
	PipelineName        string       `json:"pipelineName"`
	PipelineDescription string       `json:"pipelineDescription"`
	ProductName         string       `json:"productName"`
	ProductDescription  string       `json:"productDescription"`
	Permissions         []Permission `json:"permissions"`
}
