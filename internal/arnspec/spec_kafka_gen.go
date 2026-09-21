// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kafka
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kafka/kafka.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kafka_channel", Service: "kafka", Resource: "channel", Template: "arn:${Partition}:kafka:${Region}:${Account}:channel/${ClusterName}/${ClusterUuid}/${ChannelName}/${Uuid}"},
		{Name: "kafka_cluster", Service: "kafka", Resource: "cluster", Template: "arn:${Partition}:kafka:${Region}:${Account}:cluster/${ClusterName}/${Uuid}"},
		{Name: "kafka_configuration", Service: "kafka", Resource: "configuration", Template: "arn:${Partition}:kafka:${Region}:${Account}:configuration/${ConfigurationName}/${Uuid}"},
		{Name: "kafka_group", Service: "kafka", Resource: "group", Template: "arn:${Partition}:kafka:${Region}:${Account}:group/${ClusterName}/${ClusterUuid}/${GroupName}"},
		{Name: "kafka_replicator", Service: "kafka", Resource: "replicator", Template: "arn:${Partition}:kafka:${Region}:${Account}:replicator/${ReplicatorName}/${Uuid}"},
		{Name: "kafka_topic", Service: "kafka", Resource: "topic", Template: "arn:${Partition}:kafka:${Region}:${Account}:topic/${ClusterName}/${ClusterUuid}/${TopicName}"},
		{Name: "kafka_transactional_id", Service: "kafka", Resource: "transactional-id", Template: "arn:${Partition}:kafka:${Region}:${Account}:transactional-id/${ClusterName}/${ClusterUuid}/${TransactionalId}"},
		{Name: "kafka_vpc_connection", Service: "kafka", Resource: "vpc-connection", Template: "arn:${Partition}:kafka:${Region}:${VpcOwnerAccount}:vpc-connection/${ClusterOwnerAccount}/${ClusterName}/${Uuid}"},
	})
}
