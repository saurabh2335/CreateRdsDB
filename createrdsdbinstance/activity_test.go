package createrdsdbinstance

import (
	"io/ioutil"
	"testing"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("accessKey", "AKIAI3VJTISFSV55MY7Q")
	tc.SetInput("secretKey", "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")
	tc.SetInput("region", "ap-southeast-2")
	tc.SetInput("dbInstanceIdentifier", "flogordsinstance-test1")
	tc.SetInput("dbInstanceClass", "db.t2.micro")
	tc.SetInput("engine", "oracle-ee")
	//tc.SetInput("engineVersion", "11.2.0.4.v9")
	tc.SetInput("masterUsername", "flogordsinstancetest1")
	tc.SetInput("masterUserPassword", "flogordsinstancetest1")
	/*tc.SetInput("licenseModel", "bring-your-own-license")
	tc.SetInput("allocatedStorage", 20)
	tc.SetInput("storageType", "standard")
	tc.SetInput("port", "1250")
	tc.SetInput("availabilityZone", "ap-south-1")
	*/

	success, err := act.Eval(tc)

	if err != nil {
		t.Error("Error while Craeting the instance")
		t.Fail()
		return
	}
	if success {
		response := tc.GetOutput("response")
		status := tc.GetOutput("status")
		fmt.Printf("Response : %v\n", response)
		fmt.Printf("status: %v\n", status)
	} else {
		t.Error("Unknown Error")
		t.Fail()
		return
	}
}
