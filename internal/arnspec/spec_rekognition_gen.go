// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rekognition
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rekognition/rekognition.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rekognition_collection", Service: "rekognition", Resource: "collection", Template: "arn:${Partition}:rekognition:${Region}:${Account}:collection/${CollectionId}"},
		{Name: "rekognition_dataset", Service: "rekognition", Resource: "dataset", Template: "arn:${Partition}:rekognition:${Region}:${Account}:project/${ProjectName}/dataset/${DatasetType}/${CreationTimestamp}"},
		{Name: "rekognition_project", Service: "rekognition", Resource: "project", Template: "arn:${Partition}:rekognition:${Region}:${Account}:project/${ProjectName}/${CreationTimestamp}"},
		{Name: "rekognition_projectversion", Service: "rekognition", Resource: "projectversion", Template: "arn:${Partition}:rekognition:${Region}:${Account}:project/${ProjectName}/version/${VersionName}/${CreationTimestamp}"},
		{Name: "rekognition_streamprocessor", Service: "rekognition", Resource: "streamprocessor", Template: "arn:${Partition}:rekognition:${Region}:${Account}:streamprocessor/${StreamprocessorId}"},
	})
}
