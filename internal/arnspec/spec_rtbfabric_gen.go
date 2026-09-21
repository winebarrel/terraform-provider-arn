// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rtbfabric
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rtbfabric/rtbfabric.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rtbfabric_inbound_external_link", Service: "rtbfabric", Resource: "InboundExternalLink", Template: "arn:${Partition}:rtbfabric:${Region}:${Account}:gateway/${GatewayId}/link/${LinkId}"},
		{Name: "rtbfabric_link", Service: "rtbfabric", Resource: "Link", Template: "arn:${Partition}:rtbfabric:${Region}:${Account}:gateway/${GatewayId}/link/${LinkId}"},
		{Name: "rtbfabric_link_routing_rule", Service: "rtbfabric", Resource: "LinkRoutingRule", Template: "arn:${Partition}:rtbfabric:${Region}:${Account}:gateway/${GatewayId}/link/${LinkId}/routing-rule/${RuleId}"},
		{Name: "rtbfabric_outbound_external_link", Service: "rtbfabric", Resource: "OutboundExternalLink", Template: "arn:${Partition}:rtbfabric:${Region}:${Account}:gateway/${GatewayId}/link/${LinkId}"},
		{Name: "rtbfabric_requester_gateway", Service: "rtbfabric", Resource: "RequesterGateway", Template: "arn:${Partition}:rtbfabric:${Region}:${Account}:gateway/${GatewayId}"},
		{Name: "rtbfabric_responder_gateway", Service: "rtbfabric", Resource: "ResponderGateway", Template: "arn:${Partition}:rtbfabric:${Region}:${Account}:gateway/${GatewayId}"},
	})
}
