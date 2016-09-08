package types

type PipelineStatus struct {
	PipelineId            string `json:"pipelineId"`
	RequirementManagement string `json:"ca"`
	SoftwareControlManage string `json:"scm"`
	ContinuousIntegration string `json:"ci"`
	CodeQualityInspection string `json:"cq"`
	RepositoryForArtifact string `json:"rpa"`
	RepositoryOfContainer string `json:"rpd"`
	PipelineDashboard     string `json:"cov"`
	ContainerManagement   string `json:"cmp"`
}
