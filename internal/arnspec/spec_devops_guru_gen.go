// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: devops-guru
// Source: https://servicereference.us-east-1.amazonaws.com/v1/devops-guru/devops-guru.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "devops_guru_topic", Service: "devops-guru", Resource: "topic", Template: "arn:${Partition}:sns:${Region}:${Account}:${TopicName}"},
	})
}
