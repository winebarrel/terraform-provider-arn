// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: groundstation
// Source: https://servicereference.us-east-1.amazonaws.com/v1/groundstation/groundstation.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "groundstation_agent", Service: "groundstation", Resource: "Agent", Template: "arn:${Partition}:groundstation:${Region}:${Account}:agent/${AgentId}"},
		{Name: "groundstation_config", Service: "groundstation", Resource: "Config", Template: "arn:${Partition}:groundstation:${Region}:${Account}:config/${ConfigType}/${ConfigId}"},
		{Name: "groundstation_contact", Service: "groundstation", Resource: "Contact", Template: "arn:${Partition}:groundstation:${Region}:${Account}:contact/${ContactId}"},
		{Name: "groundstation_dataflow_endpoint_group", Service: "groundstation", Resource: "DataflowEndpointGroup", Template: "arn:${Partition}:groundstation:${Region}:${Account}:dataflow-endpoint-group/${DataflowEndpointGroupId}"},
		{Name: "groundstation_ephemeris_item", Service: "groundstation", Resource: "EphemerisItem", Template: "arn:${Partition}:groundstation:${Region}:${Account}:ephemeris/${EphemerisId}"},
		{Name: "groundstation_ground_station_resource", Service: "groundstation", Resource: "GroundStationResource", Template: "arn:${Partition}:groundstation:${Region}:${Account}:groundstation:${GroundStationId}"},
		{Name: "groundstation_mission_profile", Service: "groundstation", Resource: "MissionProfile", Template: "arn:${Partition}:groundstation:${Region}:${Account}:mission-profile/${MissionProfileId}"},
		{Name: "groundstation_satellite", Service: "groundstation", Resource: "Satellite", Template: "arn:${Partition}:groundstation:${Region}:${Account}:satellite/${SatelliteId}"},
	})
}
