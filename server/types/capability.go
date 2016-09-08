package types

type PipelineCapability struct {
	Kind        string      `json:"kind"`
	Driver      string      `json:"driver"`
	Provider    string      `json:"provider"`
	ConfigItems ConfigItems `json:"configItems"`
}

type PipelineCapabilities []PipelineCapability
