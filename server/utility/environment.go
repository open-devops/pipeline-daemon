package utility

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	"strings"
)

const (
	DEFAULT_ENGINE_PATH = "/opt/opendevops/pipelines"
	PIPELINE_ENGINE_NAM = "pipeline"
	PATH_SEPARATOR      = "/"
)

func GetEnginePath(pipelineInfo *types.PipelineInfo) string {
	return strings.Join(
		[]string{
			DEFAULT_ENGINE_PATH,
			pipelineInfo.ProductName,
			pipelineInfo.PipelineName,
		},
		PATH_SEPARATOR)
}

func GetEngineName() string {
	return PIPELINE_ENGINE_NAM
}
