// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mq
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mq/mq.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mq_brokers", Service: "mq", Resource: "brokers", Template: "arn:${Partition}:mq:${Region}:${Account}:broker:${BrokerName}:${BrokerId}"},
		{Name: "mq_configurations", Service: "mq", Resource: "configurations", Template: "arn:${Partition}:mq:${Region}:${Account}:configuration:${ConfigurationId}"},
	})
}
