package kafkaActivity

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("tibco-kafka")

	if act == nil {
		t.Error("Activity Not Registered")
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

	md := activity.NewMetadata(jsonMetadata)
	act := &KafkaActivity{metadata: md}

	tc := test.NewTestActivityContext(md)

	//setup attrs
	tc.SetInput(topic, "derrick-001")
	tc.SetInput(message, "Test Sent to NBCU Kafka")

	act.Eval(tc)

	//check result attr

	// TODO how to do some checks if the activity has no Output?
}
