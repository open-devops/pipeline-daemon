package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"os/exec"
)

func CreateProvision(pipelineInfo *types.PipelineInfo) error {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Engine program name
	engineProgramPath := utl.GetEngineProgramPath(pipelineInfo)

	// Remove existing environment
	existed, err := utl.Exists(engineParentPath)
	if err != nil {
		return err
	}
	if existed {
		if err := os.RemoveAll(engineParentPath); err != nil {
			return err
		}
	}

	// Create new environment
	if err := os.MkdirAll(engineParentPath, os.ModePerm); err != nil {
		return err
	} else {
		if err := os.Chdir(engineParentPath); err != nil {
			return err
		}
	}

	// Make engine program ready
	if err := utl.CopyFolder(utl.GetEngineTemplatePath(), engineParentPath); err != nil {
		return err
	}

	// Make engine program executable
	if err := os.Chmod(engineProgramPath, os.ModePerm); err != nil {
		return err
	}

	// Start provision process
	args := []string{"init"}
	if out, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		return err
	} else {
		fmt.Println(string(out))
		return nil
	}

	return nil
}

func DeleteProvision(pipelineInfo *types.PipelineInfo) error {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Remove existing environment
	existed, err := utl.Exists(engineParentPath)
	if err != nil {
		return err
	}
	if existed {
		if _, err := HandlePipelineAction(pipelineInfo, ACTION_STOP, nil); err != nil {
			return err
		}
		if err := os.RemoveAll(engineParentPath); err != nil {
			return err
		}
	}

	return nil
}

func FetchProvisionInfo(pipelineInfo *types.PipelineInfo) (*types.PipelineProvision, error) {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Remove existing environment
	existed, err := utl.Exists(engineParentPath)
	if err != nil {
		return &types.PipelineProvision{PipelineId: pipelineInfo.PipelineId, Capabilities: nil}, err
	}
	if !existed {
		return &types.PipelineProvision{PipelineId: pipelineInfo.PipelineId, Capabilities: nil}, nil
	}

	// At this time, because we have not retrieved provision info from the pipeline runtime
	// there we return a default supported pipeline provision just for placeholder to improve.
	return &types.PipelineProvision{
		PipelineId: pipelineInfo.PipelineId,
		Capabilities: types.PipelineCapabilities{
			types.PipelineCapability{
				Kind:        "ca",
				Driver:      "docker",
				Provider:    "Jira",
				ConfigItems: nil,
			},
			types.PipelineCapability{
				Kind:     "scm",
				Driver:   "docker",
				Provider: "Gitlab",
				ConfigItems: types.ConfigItems{
					types.ConfigItem{
						Kind:  "text",
						Name:  "Repository Name",
						Value: pipelineInfo.PipelineName,
					},
					types.ConfigItem{
						Kind:  "option",
						Name:  "Branch",
						Value: "Master",
					},
				},
			},
			types.PipelineCapability{
				Kind:        "ci",
				Driver:      "docker",
				Provider:    "Jenkins",
				ConfigItems: nil,
			},
			types.PipelineCapability{
				Kind:        "cq",
				Driver:      "docker",
				Provider:    "SonarQube",
				ConfigItems: nil,
			},
			types.PipelineCapability{
				Kind:        "rpa",
				Driver:      "docker",
				Provider:    "nexus",
				ConfigItems: nil,
			},
			types.PipelineCapability{
				Kind:        "rpd",
				Driver:      "docker",
				Provider:    "Harbor",
				ConfigItems: nil,
			},
			types.PipelineCapability{
				Kind:        "cov",
				Driver:      "docker",
				Provider:    "Hygieia",
				ConfigItems: nil,
			},
			types.PipelineCapability{
				Kind:        "cmp",
				Driver:      "docker",
				Provider:    "Rancher",
				ConfigItems: nil,
			},
		},
	}, nil
}
