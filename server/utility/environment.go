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

func GetEngineParentPath(pipelineInfo *types.PipelineInfo) string {
	return strings.Join(
		[]string{
			DEFAULT_ENGINE_PATH,
			pipelineInfo.ProductName,
			pipelineInfo.PipelineName,
		},
		PATH_SEPARATOR)
}

func GetEngineProgramPath(pipelineInfo *types.PipelineInfo) string {
	return strings.Join(
		[]string{
			DEFAULT_ENGINE_PATH,
			pipelineInfo.ProductName,
			pipelineInfo.PipelineName,
			PATH_SEPARATOR,
		},
		PATH_SEPARATOR)
}
