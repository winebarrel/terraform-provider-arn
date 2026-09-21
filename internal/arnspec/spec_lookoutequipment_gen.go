// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: lookoutequipment
// Source: https://servicereference.us-east-1.amazonaws.com/v1/lookoutequipment/lookoutequipment.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "lookoutequipment_dataset", Service: "lookoutequipment", Resource: "dataset", Template: "arn:${Partition}:lookoutequipment:${Region}:${Account}:dataset/${DatasetName}/${DatasetId}"},
		{Name: "lookoutequipment_inference_scheduler", Service: "lookoutequipment", Resource: "inference-scheduler", Template: "arn:${Partition}:lookoutequipment:${Region}:${Account}:inference-scheduler/${InferenceSchedulerName}/${InferenceSchedulerId}"},
		{Name: "lookoutequipment_label_group", Service: "lookoutequipment", Resource: "label-group", Template: "arn:${Partition}:lookoutequipment:${Region}:${Account}:label-group/${LabelGroupName}/${LabelGroupId}"},
		{Name: "lookoutequipment_model", Service: "lookoutequipment", Resource: "model", Template: "arn:${Partition}:lookoutequipment:${Region}:${Account}:model/${ModelName}/${ModelId}"},
		{Name: "lookoutequipment_model_version", Service: "lookoutequipment", Resource: "model-version", Template: "arn:${Partition}:lookoutequipment:${Region}:${Account}:model/${ModelName}/${ModelId}/model-version/${ModelVersionNumber}"},
	})
}
