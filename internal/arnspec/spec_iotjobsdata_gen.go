// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotjobsdata
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotjobsdata/iotjobsdata.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotjobsdata_thing", Service: "iotjobsdata", Resource: "thing", Template: "arn:${Partition}:iot:${Region}:${Account}:thing/${ThingName}"},
	})
}
