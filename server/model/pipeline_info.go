package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
)

func FetchPipelineInfo(pipelineId string) *types.PipelineInfo {
	pipelineInfo := &types.PipelineInfo{
		PipelineId: pipelineId,
		PipelineName: "OpenDevOps",
	}

	return pipelineInfo
}
