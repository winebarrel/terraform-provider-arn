// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotevents
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotevents/iotevents.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotevents_alarm_model", Service: "iotevents", Resource: "alarmModel", Template: "arn:${Partition}:iotevents:${Region}:${Account}:alarmModel/${AlarmModelName}"},
		{Name: "iotevents_detector_model", Service: "iotevents", Resource: "detectorModel", Template: "arn:${Partition}:iotevents:${Region}:${Account}:detectorModel/${DetectorModelName}"},
		{Name: "iotevents_input", Service: "iotevents", Resource: "input", Template: "arn:${Partition}:iotevents:${Region}:${Account}:input/${InputName}"},
	})
}
