// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kafka-cluster
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kafka-cluster/kafka-cluster.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kafka_cluster_cluster", Service: "kafka-cluster", Resource: "cluster", Template: "arn:${Partition}:kafka:${Region}:${Account}:cluster/${ClusterName}/${ClusterUuid}"},
		{Name: "kafka_cluster_group", Service: "kafka-cluster", Resource: "group", Template: "arn:${Partition}:kafka:${Region}:${Account}:group/${ClusterName}/${ClusterUuid}/${GroupName}"},
		{Name: "kafka_cluster_topic", Service: "kafka-cluster", Resource: "topic", Template: "arn:${Partition}:kafka:${Region}:${Account}:topic/${ClusterName}/${ClusterUuid}/${TopicName}"},
		{Name: "kafka_cluster_transactional_id", Service: "kafka-cluster", Resource: "transactional-id", Template: "arn:${Partition}:kafka:${Region}:${Account}:transactional-id/${ClusterName}/${ClusterUuid}/${TransactionalId}"},
	})
}
