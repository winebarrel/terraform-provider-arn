// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: networkmanager
// Source: https://servicereference.us-east-1.amazonaws.com/v1/networkmanager/networkmanager.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "networkmanager_attachment", Service: "networkmanager", Resource: "attachment", Template: "arn:${Partition}:networkmanager::${Account}:attachment/${ResourceId}"},
		{Name: "networkmanager_connect_peer", Service: "networkmanager", Resource: "connect-peer", Template: "arn:${Partition}:networkmanager::${Account}:connect-peer/${ResourceId}"},
		{Name: "networkmanager_connection", Service: "networkmanager", Resource: "connection", Template: "arn:${Partition}:networkmanager::${Account}:connection/${GlobalNetworkId}/${ResourceId}"},
		{Name: "networkmanager_core_network", Service: "networkmanager", Resource: "core-network", Template: "arn:${Partition}:networkmanager::${Account}:core-network/${ResourceId}"},
		{Name: "networkmanager_device", Service: "networkmanager", Resource: "device", Template: "arn:${Partition}:networkmanager::${Account}:device/${GlobalNetworkId}/${ResourceId}"},
		{Name: "networkmanager_global_network", Service: "networkmanager", Resource: "global-network", Template: "arn:${Partition}:networkmanager::${Account}:global-network/${ResourceId}"},
		{Name: "networkmanager_link", Service: "networkmanager", Resource: "link", Template: "arn:${Partition}:networkmanager::${Account}:link/${GlobalNetworkId}/${ResourceId}"},
		{Name: "networkmanager_peering", Service: "networkmanager", Resource: "peering", Template: "arn:${Partition}:networkmanager::${Account}:peering/${ResourceId}"},
		{Name: "networkmanager_site", Service: "networkmanager", Resource: "site", Template: "arn:${Partition}:networkmanager::${Account}:site/${GlobalNetworkId}/${ResourceId}"},
	})
}
