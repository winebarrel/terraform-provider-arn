// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: medialive
// Source: https://servicereference.us-east-1.amazonaws.com/v1/medialive/medialive.json
// Functions: 17
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "medialive_channel", Service: "medialive", Resource: "channel", Template: "arn:${Partition}:medialive:${Region}:${Account}:channel:${ChannelId}"},
		{Name: "medialive_channel_placement_group", Service: "medialive", Resource: "channel-placement-group", Template: "arn:${Partition}:medialive:${Region}:${Account}:channelPlacementGroup:${ClusterId}/${ChannelPlacementGroupId}"},
		{Name: "medialive_cloudwatch_alarm_template", Service: "medialive", Resource: "cloudwatch-alarm-template", Template: "arn:${Partition}:medialive:${Region}:${Account}:cloudwatch-alarm-template:${CloudWatchAlarmTemplateId}"},
		{Name: "medialive_cloudwatch_alarm_template_group", Service: "medialive", Resource: "cloudwatch-alarm-template-group", Template: "arn:${Partition}:medialive:${Region}:${Account}:cloudwatch-alarm-template-group:${CloudWatchAlarmTemplateGroupId}"},
		{Name: "medialive_cluster", Service: "medialive", Resource: "cluster", Template: "arn:${Partition}:medialive:${Region}:${Account}:cluster:${ClusterId}"},
		{Name: "medialive_eventbridge_rule_template", Service: "medialive", Resource: "eventbridge-rule-template", Template: "arn:${Partition}:medialive:${Region}:${Account}:eventbridge-rule-template:${EventBridgeRuleTemplateId}"},
		{Name: "medialive_eventbridge_rule_template_group", Service: "medialive", Resource: "eventbridge-rule-template-group", Template: "arn:${Partition}:medialive:${Region}:${Account}:eventbridge-rule-template-group:${EventBridgeRuleTemplateGroupId}"},
		{Name: "medialive_input", Service: "medialive", Resource: "input", Template: "arn:${Partition}:medialive:${Region}:${Account}:input:${InputId}"},
		{Name: "medialive_input_device", Service: "medialive", Resource: "input-device", Template: "arn:${Partition}:medialive:${Region}:${Account}:inputDevice:${DeviceId}"},
		{Name: "medialive_input_security_group", Service: "medialive", Resource: "input-security-group", Template: "arn:${Partition}:medialive:${Region}:${Account}:inputSecurityGroup:${InputSecurityGroupId}"},
		{Name: "medialive_multiplex", Service: "medialive", Resource: "multiplex", Template: "arn:${Partition}:medialive:${Region}:${Account}:multiplex:${MultiplexId}"},
		{Name: "medialive_network", Service: "medialive", Resource: "network", Template: "arn:${Partition}:medialive:${Region}:${Account}:network:${NetworkId}"},
		{Name: "medialive_node", Service: "medialive", Resource: "node", Template: "arn:${Partition}:medialive:${Region}:${Account}:node:${ClusterId}/${NodeId}"},
		{Name: "medialive_offering", Service: "medialive", Resource: "offering", Template: "arn:${Partition}:medialive:${Region}:${Account}:offering:${OfferingId}"},
		{Name: "medialive_reservation", Service: "medialive", Resource: "reservation", Template: "arn:${Partition}:medialive:${Region}:${Account}:reservation:${ReservationId}"},
		{Name: "medialive_sdi_source", Service: "medialive", Resource: "sdi-source", Template: "arn:${Partition}:medialive:${Region}:${Account}:sdiSource:${SdiSourceId}"},
		{Name: "medialive_signal_map", Service: "medialive", Resource: "signal-map", Template: "arn:${Partition}:medialive:${Region}:${Account}:signal-map:${SignalMapId}"},
	})
}
