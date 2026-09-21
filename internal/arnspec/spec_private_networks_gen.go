// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: private-networks
// Source: https://servicereference.us-east-1.amazonaws.com/v1/private-networks/private-networks.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "private_networks_device_identifier", Service: "private-networks", Resource: "device-identifier", Template: "arn:${Partition}:private-networks:${Region}:${Account}:device-identifier/${NetworkName}/${DeviceId}"},
		{Name: "private_networks_network", Service: "private-networks", Resource: "network", Template: "arn:${Partition}:private-networks:${Region}:${Account}:network/${NetworkName}"},
		{Name: "private_networks_network_resource", Service: "private-networks", Resource: "network-resource", Template: "arn:${Partition}:private-networks:${Region}:${Account}:network-resource/${NetworkName}/${ResourceId}"},
		{Name: "private_networks_network_site", Service: "private-networks", Resource: "network-site", Template: "arn:${Partition}:private-networks:${Region}:${Account}:network-site/${NetworkName}/${NetworkSiteName}"},
		{Name: "private_networks_order", Service: "private-networks", Resource: "order", Template: "arn:${Partition}:private-networks:${Region}:${Account}:order/${NetworkName}/${OrderId}"},
	})
}
