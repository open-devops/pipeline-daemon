package types

type PipelineCapability struct {
	kind        string
	driver      string
	provider    string
	configItems []PipelineCapabilityConfigItem
}
