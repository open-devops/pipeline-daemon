package model

import (
	"encoding/json"
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type Pipeline struct {
	ID                  bson.ObjectId `_id,omitempty`
	Name                string        `name,omitempty`
	Description         string        `description,omitempty`
	PipelineId          string        `pipelineId,omitempty`
	PipelineName        string        `pipelineName,omitempty`
	PipelineDescription string        `pipelineDescription,omitempty`
	OrganizationId      string        `organizationId,omitempty`
}

type Account struct {
	ID          bson.ObjectId `_id,omitempty`
	Name        string        `name,omitempty`
	Mail        string        `mail,omitempty`
	AccessToken string        `accessToken,omitempty`
}

type Role struct {
	ID   bson.ObjectId `_id,omitempty`
	Name string        `name,omitempty`
}

func FetchPipelineInfo(pipelineId string) *types.PipelineInfo {
	// Create session
	session, err := mgo.Dial(os.Getenv("PORTAL_LOCALHOST"))
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

	// Fetch Pipeline permission information
	var permissions []types.Permission
	c = session.DB("dashboard").C("permission")
	err = c.Find(bson.M{"pipelineId": pipelineId}).All(&permissions)
	if err != nil {
		return new(types.PipelineInfo) // Pipeline not exist
	}

	// Fetch Account information
	c = session.DB("dashboard").C("account")
	for i := 0; i < len(permissions); i++ {
		permission := &permissions[i]
		account := Account{}
		err = c.Find(bson.M{"_id": permission.AccountId}).One(&account)
		permission.AccountName = account.Name
		permission.AccountMail = account.Mail
		permission.AccessToken = account.AccessToken
	}

	// Fetch Role information
	c = session.DB("dashboard").C("role")
	for i := 0; i < len(permissions); i++ {
		permission := &permissions[i]
		role := Role{}
		err = c.Find(bson.M{"_id": permission.RoleId}).One(&role)
		permission.RoleName = role.Name
	}

	// Return Pipeline Information
	pipelineInfo := &types.PipelineInfo{
		PipelineId:          pipeline.PipelineId,
		PipelineName:        pipeline.PipelineName,
		PipelineDescription: pipeline.PipelineDescription,
		ProductName:         pipeline.Name,
		ProductDescription:  pipeline.Description,
		Permissions:         permissions,
	}

	// TODO: delete this debug information
	debugInfo, _ := json.Marshal(pipelineInfo)
	fmt.Println(string(debugInfo))

	return pipelineInfo
}
