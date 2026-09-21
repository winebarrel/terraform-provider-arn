// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotwireless
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotwireless/iotwireless.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotwireless_cert", Service: "iotwireless", Resource: "cert", Template: "arn:${Partition}:iot:${Region}:${Account}:cert/${Certificate}"},
		{Name: "iotwireless_destination", Service: "iotwireless", Resource: "Destination", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:Destination/${DestinationName}"},
		{Name: "iotwireless_device_profile", Service: "iotwireless", Resource: "DeviceProfile", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:DeviceProfile/${DeviceProfileId}"},
		{Name: "iotwireless_fuota_task", Service: "iotwireless", Resource: "FuotaTask", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:FuotaTask/${FuotaTaskId}"},
		{Name: "iotwireless_import_task", Service: "iotwireless", Resource: "ImportTask", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:ImportTask/${ImportTaskId}"},
		{Name: "iotwireless_multicast_group", Service: "iotwireless", Resource: "MulticastGroup", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:MulticastGroup/${MulticastGroupId}"},
		{Name: "iotwireless_network_analyzer_configuration", Service: "iotwireless", Resource: "NetworkAnalyzerConfiguration", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:NetworkAnalyzerConfiguration/${NetworkAnalyzerConfigurationName}"},
		{Name: "iotwireless_service_profile", Service: "iotwireless", Resource: "ServiceProfile", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:ServiceProfile/${ServiceProfileId}"},
		{Name: "iotwireless_sidewalk_account", Service: "iotwireless", Resource: "SidewalkAccount", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:SidewalkAccount/${SidewalkAccountId}"},
		{Name: "iotwireless_thing", Service: "iotwireless", Resource: "thing", Template: "arn:${Partition}:iot:${Region}:${Account}:thing/${ThingName}"},
		{Name: "iotwireless_wireless_device", Service: "iotwireless", Resource: "WirelessDevice", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:WirelessDevice/${WirelessDeviceId}"},
		{Name: "iotwireless_wireless_gateway", Service: "iotwireless", Resource: "WirelessGateway", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:WirelessGateway/${WirelessGatewayId}"},
		{Name: "iotwireless_wireless_gateway_task_definition", Service: "iotwireless", Resource: "WirelessGatewayTaskDefinition", Template: "arn:${Partition}:iotwireless:${Region}:${Account}:WirelessGatewayTaskDefinition/${WirelessGatewayTaskDefinitionId}"},
	})
}
