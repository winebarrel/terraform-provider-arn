// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotfleetwise
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotfleetwise/iotfleetwise.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotfleetwise_campaign", Service: "iotfleetwise", Resource: "campaign", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:campaign/${CampaignName}"},
		{Name: "iotfleetwise_decodermanifest", Service: "iotfleetwise", Resource: "decodermanifest", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:decoder-manifest/${Name}"},
		{Name: "iotfleetwise_fleet", Service: "iotfleetwise", Resource: "fleet", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:fleet/${FleetId}"},
		{Name: "iotfleetwise_modelmanifest", Service: "iotfleetwise", Resource: "modelmanifest", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:model-manifest/${Name}"},
		{Name: "iotfleetwise_signalcatalog", Service: "iotfleetwise", Resource: "signalcatalog", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:signal-catalog/${Name}"},
		{Name: "iotfleetwise_statetemplate", Service: "iotfleetwise", Resource: "statetemplate", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:state-template/${StateTemplateId}"},
		{Name: "iotfleetwise_vehicle", Service: "iotfleetwise", Resource: "vehicle", Template: "arn:${Partition}:iotfleetwise:${Region}:${Account}:vehicle/${VehicleId}"},
	})
}
