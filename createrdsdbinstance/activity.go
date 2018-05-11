package createrdsdbinstance

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"fmt"
)

var log = logger.GetLogger("activity-tibco-createrdsdbinstance-oracle")

const (
	ivRegion    					= "region"
	ivAccessKey 					= "accessKey"
	ivSecretKey 					= "secretKey"
	ivDBInstanceIdentifier 			= "dbInstanceIdentifier"
	ivDBInstanceClass				= "dbInstanceClass"
	ivAllocatedStorage				= "allocatedStorage"
	ivAutoMinorVersionUpgrade		= "autoMinorVersionUpgrade"
	ivAvailabilityZone				= "availabilityZone"
	ivBackupRetentionPeriod 		= "backupRetentionPeriod"
	ivCharacterSetName				= "characterSetName"
	ivCopyTagsToSnapshot			= "copyTagsToSnapshot"
	ivMultiAZ						= "multiAZ"
	ivDBClusterIdentifier			= "dbClusterIdentifier"
	ivDBName						= "dbName"
	ivDBParameterGroupName			= "dbParameterGroupName"
	ivDBSubnetGroupName				= "dbSubnetGroupName"
	ivEngine						= "engine"
	ivEngineVersion					= "engineVersion"
	ivEnableIAMDatabaseAuthentication	= "enableIAMDatabaseAuthentication"
	ivEnablePerformanceInsights		= "enablePerformanceInsights"
	ivIops							= "iops"
	ivLicenseModel 					= "licenseModel"
	ivMasterUsername				= "masterUsername"
	ivMasterUserPassword			= "masterUserPassword"
	ivMonitoringInterval			= "monitoringInterval"
	ivMonitoringRoleArn				= "monitoringRoleArn"
	ivPort							= "port"
	ivPreferredBackupWindow			= "preferredBackupWindow"
	ivPreferredMaintenanceWindow	= "preferredMaintenanceWindow"
	ivPubliclyAccessible			= "publiclyAccessible"
	ivstorageEncrypted				= "storageEncrypted"
	ivStorageType					= "storageType"
	ivTdeCredentialArn				= "tdeCredentialArn"
	ivTdeCredentialPassword			= "tdeCredentialPassword"

	ovDBResponse	    		= "response"
	ovDBInstanceStatus	   		= "status"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	if len(context.GetInput(ivAccessKey).(string)) < 1 || len(context.GetInput(ivSecretKey).(string)) < 1 || len(context.GetInput(ivRegion).(string)) < 1 || len(context.GetInput(ivDBInstanceIdentifier).(string)) < 1 || len(context.GetInput(ivDBInstanceClass).(string)) < 1 || len(context.GetInput(ivEngine).(string)) < 1 || len(context.GetInput(ivMasterUsername).(string)) < 1 || len(context.GetInput(ivMasterUserPassword).(string)) < 1  {
		log.Error("Required variables have not been set !")
		return false, fmt.Errorf("required variables have not been set")
	}

	//To set the default the parameter in the input parameters are null.
	var dupAllocatedStorage,dupPort int
	var dupLicenseModel,dupStorageType,dupAvailabilityZone,dupDBName string

	if context.GetInput(ivAllocatedStorage) == 0{
		dupAllocatedStorage = 20
	}else{
		dupAllocatedStorage =  context.GetInput(ivAllocatedStorage).(int)
	}
	if len(context.GetInput(ivLicenseModel).(string)) < 1{
		dupLicenseModel = "bring-your-own-license"
	}else{
		dupLicenseModel =  context.GetInput(ivLicenseModel).(string)
	}
	if context.GetInput(ivPort) == 0{
		dupPort = 1521
	}else{
		dupPort =  context.GetInput(ivPort).(int)
	}
	if len(context.GetInput(ivStorageType).(string)) < 1{
		dupStorageType = "standard"
	}else{
		dupStorageType =  context.GetInput(ivStorageType).(string)
	}
	if len(context.GetInput(ivAvailabilityZone).(string)) < 1{
		dupAvailabilityZone = "ap-southeast-2a"
	}else{
		dupAvailabilityZone =  context.GetInput(ivAvailabilityZone).(string)
	}
	if len(context.GetInput(ivDBName).(string)) < 1{
		dupDBName = "ORCL"
	}else{
		dupDBName =  context.GetInput(ivDBName).(string)
	}

	var accessKey, secretKey = "", ""
	if context.GetInput(ivAccessKey) != nil {
		accessKey = context.GetInput(ivAccessKey).(string)
	}
	if context.GetInput(ivSecretKey) != nil {
		secretKey = context.GetInput(ivSecretKey).(string)
	}
	var config *aws.Config
	region := context.GetInput(ivRegion).(string)
	if accessKey != "" && secretKey != "" {
		config = aws.NewConfig().WithRegion(region).WithCredentials(credentials.NewStaticCredentials(accessKey, secretKey, ""))
	} else {
		config = aws.NewConfig().WithRegion(region)
	}
	log.Debug("Session created")

	svc := rds.New(session.New(config))

	log.Debug("Setting CreateDBInstanceInput parameters")

	input := &rds.CreateDBInstanceInput{
		DBInstanceIdentifier: aws.String(context.GetInput(ivDBInstanceIdentifier).(string)),
		DBInstanceClass: aws.String(context.GetInput(ivDBInstanceClass).(string)),
		AllocatedStorage: aws.Int64(int64(dupAllocatedStorage)),
		AutoMinorVersionUpgrade: aws.Bool(context.GetInput(ivAutoMinorVersionUpgrade).(bool)),
		AvailabilityZone: aws.String(dupAvailabilityZone),
		BackupRetentionPeriod: aws.Int64(int64(context.GetInput(ivBackupRetentionPeriod).(int))),
		DBName: aws.String(dupDBName),
		Engine: aws.String(context.GetInput(ivEngine).(string)),
		Iops: aws.Int64(int64(context.GetInput(ivIops).(int))),
		LicenseModel: aws.String(dupLicenseModel),
		MasterUsername: aws.String(context.GetInput(ivMasterUsername).(string)),
		MasterUserPassword: aws.String(context.GetInput(ivMasterUserPassword).(string)),
		MonitoringInterval: aws.Int64(int64(context.GetInput(ivMonitoringInterval).(int))),
		Port: aws.Int64(int64(dupPort)),
		PubliclyAccessible: aws.Bool(context.GetInput(ivPubliclyAccessible).(bool)),
		StorageEncrypted: aws.Bool(context.GetInput(ivstorageEncrypted).(bool)),
		StorageType: aws.String(dupStorageType),
		CopyTagsToSnapshot: aws.Bool(context.GetInput(ivCopyTagsToSnapshot).(bool)),
		EnableIAMDatabaseAuthentication: aws.Bool(context.GetInput(ivEnableIAMDatabaseAuthentication).(bool)),
		EnablePerformanceInsights: aws.Bool(context.GetInput(ivEnablePerformanceInsights).(bool)),
		MultiAZ: aws.Bool(context.GetInput(ivMultiAZ).(bool)),
	}

	//As values of the below parameter initially set to ""(NULL),so Need to filter the values without ""
	if len(context.GetInput(ivTdeCredentialArn).(string)) > 1 && len(context.GetInput(ivTdeCredentialPassword).(string)) > 1{
		input.TdeCredentialArn = aws.String(context.GetInput(ivTdeCredentialArn).(string))
		input.TdeCredentialPassword = aws.String(context.GetInput(ivTdeCredentialPassword).(string))
	}
	if len(context.GetInput(ivCharacterSetName).(string))  >  1{
		input.CharacterSetName = aws.String(context.GetInput(ivCharacterSetName).(string))
	}
	if len(context.GetInput(ivDBClusterIdentifier).(string))  >  1{
		input.DBClusterIdentifier = aws.String(context.GetInput(ivDBClusterIdentifier).(string))
	}
	if len(context.GetInput(ivDBParameterGroupName).(string))  >  1{
		input.DBParameterGroupName = aws.String(context.GetInput(ivDBParameterGroupName).(string))
	}
	if len(context.GetInput(ivDBSubnetGroupName).(string))  >  1 {
		input.DBSubnetGroupName = aws.String(context.GetInput(ivDBSubnetGroupName).(string))
	}
	if len(context.GetInput(ivMonitoringRoleArn).(string))  >  1{
		input.MonitoringRoleArn = aws.String(context.GetInput(ivMonitoringRoleArn).(string))
	}
	if len(context.GetInput(ivPreferredBackupWindow).(string))  >  1{
		input.PreferredBackupWindow = aws.String(context.GetInput(ivPreferredBackupWindow).(string))
	}
	if len(context.GetInput(ivPreferredMaintenanceWindow).(string))  >  1{
		input.PreferredMaintenanceWindow = aws.String(context.GetInput(ivPreferredMaintenanceWindow).(string))
	}
	if len(context.GetInput(ivEngineVersion).(string))  >  1{
	input.EngineVersion = aws.String(context.GetInput(ivEngineVersion).(string))
	}
	//

	fmt.Println("Input : \n",input)

	result, err := svc.CreateDBInstance(input)

	if err != nil {
		log.Error("Error Occured while creating the Instance :",err.Error())
		return false, err
	}

	log.Debugf("RDS DB Instance successfully created With DBInstanceIdentifier : ",result.DBInstance.DBInstanceIdentifier,"\n and Response : \n",result.DBInstance)
	context.SetOutput(ovDBResponse,result.DBInstance)
	context.SetOutput(ovDBInstanceStatus,"Success")

	return true, nil
}
