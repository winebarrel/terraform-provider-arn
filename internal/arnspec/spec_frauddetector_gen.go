// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: frauddetector
// Source: https://servicereference.us-east-1.amazonaws.com/v1/frauddetector/frauddetector.json
// Functions: 14
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "frauddetector_batch_import", Service: "frauddetector", Resource: "batch-import", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:batch-import/${ResourcePath}"},
		{Name: "frauddetector_batch_prediction", Service: "frauddetector", Resource: "batch-prediction", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:batch-prediction/${ResourcePath}"},
		{Name: "frauddetector_detector", Service: "frauddetector", Resource: "detector", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:detector/${ResourcePath}"},
		{Name: "frauddetector_detector_version", Service: "frauddetector", Resource: "detector-version", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:detector-version/${ResourcePath}"},
		{Name: "frauddetector_entity_type", Service: "frauddetector", Resource: "entity-type", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:entity-type/${ResourcePath}"},
		{Name: "frauddetector_event_type", Service: "frauddetector", Resource: "event-type", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:event-type/${ResourcePath}"},
		{Name: "frauddetector_external_model", Service: "frauddetector", Resource: "external-model", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:external-model/${ResourcePath}"},
		{Name: "frauddetector_label", Service: "frauddetector", Resource: "label", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:label/${ResourcePath}"},
		{Name: "frauddetector_list", Service: "frauddetector", Resource: "list", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:list/${ResourcePath}"},
		{Name: "frauddetector_model", Service: "frauddetector", Resource: "model", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:model/${ResourcePath}"},
		{Name: "frauddetector_model_version", Service: "frauddetector", Resource: "model-version", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:model-version/${ResourcePath}"},
		{Name: "frauddetector_outcome", Service: "frauddetector", Resource: "outcome", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:outcome/${ResourcePath}"},
		{Name: "frauddetector_rule", Service: "frauddetector", Resource: "rule", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:rule/${ResourcePath}"},
		{Name: "frauddetector_variable", Service: "frauddetector", Resource: "variable", Template: "arn:${Partition}:frauddetector:${Region}:${Account}:variable/${ResourcePath}"},
	})
}
