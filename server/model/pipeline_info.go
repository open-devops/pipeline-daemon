package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type Pipeline struct {
	ID                  bson.ObjectId `_id,omitempty`
	Name                string
	Description         string
	PipelineId          string
	PipelineName        string
	PipelineDescription string
	OrganizationId      string
}

func FetchPipelineInfo(pipelineId string) *types.PipelineInfo {
	// Create session
	session, err := mgo.Dial(os.Getenv("DOCKER_HOST"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Fetch Pipeline information
	pipeline := Pipeline{}
	c := session.DB("dashboard").C("product")
	err = c.Find(bson.M{"pipelineId": pipelineId}).One(&pipeline)
	if err != nil {
		return new(types.PipelineInfo) // Pipeline not exist
	}

	fmt.Println("Pipeline:", pipelineId+"="+pipeline.PipelineId+"-"+pipeline.Name)

	// return Pipeline Information
	pipelineInfo := &types.PipelineInfo{
		PipelineId:   pipeline.PipelineId,
		PipelineName: pipeline.Name,
	}

	return pipelineInfo
}
