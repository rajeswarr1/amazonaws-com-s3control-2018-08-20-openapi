package main

import (
	"github.com/aws-s3-control/mcp-server/config"
	"github.com/aws-s3-control/mcp-server/models"
	tools_v20180820 "github.com/aws-s3-control/mcp-server/tools/v20180820"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v20180820.CreateDeleteaccesspointpolicyforobjectlambdaTool(cfg),
		tools_v20180820.CreateDeletebucketreplicationTool(cfg),
		tools_v20180820.CreateDeletebucketTool(cfg),
		tools_v20180820.CreateDeletebucketlifecycleconfigurationTool(cfg),
		tools_v20180820.CreateDeleteaccesspointpolicyTool(cfg),
		tools_v20180820.CreateDeletebucketpolicyTool(cfg),
		tools_v20180820.CreateDeletestoragelensconfigurationTool(cfg),
		tools_v20180820.CreateDeleteaccesspointforobjectlambdaTool(cfg),
		tools_v20180820.CreateDeletepublicaccessblockTool(cfg),
		tools_v20180820.CreateDeletebuckettaggingTool(cfg),
		tools_v20180820.CreateDeleteaccesspointTool(cfg),
	}
}
