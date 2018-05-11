---
title: Create AWS RDS Oracle Instance
---

# Trigger Create RDS Oracle Instance function
This activity allows you to Create AWS RDS via DBInstanceIdentifier,DBInstanceClass,Engine and provide the access key and secret for authentication.

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "accessKey",
      "type": "string",
      "required": "true"
    },
    {
      "name": "secretKey",
      "type": "string",
      "required": "true"
    },
    {
      "name": "region",
      "type": "string",
      "required": "true",
      "allowed" : ["us-east-2","us-east-1","us-west-1","us-west-2","ap-south-1","ap-northeast-2","ap-southeast-1","ap-southeast-2","ap-northeast-1","cn-northwest-1","ca-central-1","eu-central-1","eu-west-1","eu-west-2","sa-east-1"]
    },
	{
      "name": "dbInstanceIdentifier",
      "type": "string",
      "required": "true"
    },
	{
      "name": "dbInstanceClass",
      "type": "string",
      "required": "true",
	  "allowed" : ["db.t2.micro","db.t2.small","db.t2.medium","db.t2.large","db.t2.xlarge","db.t2.2xlarge","db.m2.xlarge","db.m2.2xlarge","db.m2.4xlarge","db.r3.large","db.r3.xlarge","db.r3.2xlarge","db.r3.4xlarge","db.r3.8xlarge","db.r4.large","db.r4.xlarge","db.r4.2xlarge","db.r4.4xlarge","db.r4.8xlarge","db.r4.16xlarge","db.m1.small","db.m1.medium","db.m1.large","db.m1.xlarge","db.m3.medium","db.m3.large","db.m3.xlarge","db.m3.2xlarge","db.m4.large","db.m4.xlarge","db.m4.2xlarge","db.m4.4xlarge","db.m4.10xlarge","db.m4.16xlarge"]
    },
	{
      "name": "allocatedStorage",
      "type": "integer",
	  "minimum": 20,
	  "maximum": 16384,
      "required": "true"
    },
	{
      "name": "autoMinorVersionUpgrade",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "availabilityZone",
      "type": "string",
	  "allowed" : ["us-east-2","us-east-1","us-west-1","us-west-2","ap-south-1","ap-northeast-2","ap-southeast-1","ap-southeast-2","ap-northeast-1","cn-northwest-1","ca-central-1","eu-central-1","eu-west-1","eu-west-2","sa-east-1"],
	  "required": "false"
    },
	{
      "name": "backupRetentionPeriod",
      "type": "integer",
	  "minimum": 0,
	  "maximum": 35,
      "required": "false"

    },
	{
      "name": "characterSetName",
      "type": "string",
      "required": "false"
    },
	{
      "name": "copyTagsToSnapshot",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "multiAZ",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "dbClusterIdentifier",
      "type": "string",
      "required": "false"
    },
	{
      "name": "dbName",
      "type": "string",
	  "maxLength": 35,
      "required": "false"
    },
	{
      "name": "dbParameterGroupName",
      "type": "string",
	  "minLength": 0,
	  "maxLength": 255,
      "required": "false"
    },
	{
      "name": "dbSubnetGroupName",
      "type": "string",
      "required": "false"
    },
	{
      "name": "engine",
      "type": "string",
	  "allowed" : ["oracle-ee","oracle-se2","oracle-se1","oracle-se"],
      "required": "true"
    },
	{
      "name": "engineVersion",
      "type": "string",
	  "allowed" : ["12.1.0.2.v1","12.1.0.2.v2","12.1.0.2.v3","12.1.0.2.v4","12.1.0.2.v5","12.1.0.2.v6","12.1.0.2.v7","12.1.0.2.v8","12.1.0.2.v9","11.2.0.4.v1","11.2.0.4.v2","11.2.0.4.v3","11.2.0.4.v4","11.2.0.4.v5","11.2.0.4.v6","11.2.0.4.v7","11.2.0.4.v8","11.2.0.4.v9","11.2.0.4.v10","11.2.0.4.v11","11.2.0.4.v12","11.2.0.4.v13"]
    },
	{
      "name": "enableIAMDatabaseAuthentication",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "enablePerformanceInsights",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "iops",
      "type": "integer",
      "required": "false"
    },
	{
      "name": "licenseModel",
      "type": "string",
	  "allowed" : ["license-included","bring-your-own-license","general-public-license"]
    },
	{
      "name": "masterUsername",
      "type": "string",
	  "minLength": 1,
	  "maxLength": 30,
      "required": "true"
    },
	{
      "name": "masterUserPassword",
      "type": "string",
	  "minLength": 8,
	  "maxLength": 40,
      "required": "true"
    },
	{
      "name": "monitoringInterval",
      "type": "integer",
	  "allowed" : [0, 1, 5, 10, 15, 30, 60]
    },
	{
      "name": "monitoringRoleArn",
      "type": "string",
      "required": "false"
    },
	{
      "name": "port",
      "type": "integer",
	  "minimum": 1150,
	  "maximum": 65530
    },
	{
      "name": "preferredBackupWindow",
      "type": "string",
      "required": "false"
    },
	{
      "name": "preferredMaintenanceWindow",
      "type": "string",
      "required": "false"
    },
	{
      "name": "publiclyAccessible",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "storageEncrypted",
      "type": "boolean",
      "required": "false"
    },
	{
      "name": "storageType",
      "type": "string",
	  "allowed" : ["standard","gp2","io1"]
    },
	{
      "name": "tdeCredentialArn",
      "type": "string",
      "required": "false"
    },
	{
      "name": "tdeCredentialPassword",
      "type": "string",
      "required": "false"
    }
  ],
  "outputs": [
	{
      "name": "response",
      "type": "string"
    },
	{
      "name": "status",
      "type": "string"
    }
  ]
}
```

## Settings
| Setting     				| Required | Description 											|
|:----------------------	|:---------|:---------------------------------------------------------------------------------|
| region      				| True     | The AWS region in which you want to invoke the function |
| accessKey   				| True     | AWS access key for the user to invoke the function |
| secretKey   				| True     | AWS secret key for the user to invoke te function |
| dbInstanceIdentifier  	| True     | String value dbInstanceIdentifier must be existed as AWS Instance|
| dbInstanceClass			| True     | String value finalDBSnapshotIdentifier for finalDBSnapshotIdentifier. |
| allocatedStorage  		| False    | A boolean value for skipFinalSnapshot. |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| autoMinorVersionUpgrade   | False    | A dbInstanceIdentifier from response    |
| response      			| False    | The response from the invocation |
| status      				| False    | The status of the invocation |

## Examples
Coming soon...