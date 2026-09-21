// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: comprehend
// Source: https://servicereference.us-east-1.amazonaws.com/v1/comprehend/comprehend.json
// Functions: 15
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "comprehend_document_classification_job", Service: "comprehend", Resource: "document-classification-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:document-classification-job/${JobId}"},
		{Name: "comprehend_document_classifier", Service: "comprehend", Resource: "document-classifier", Template: "arn:${Partition}:comprehend:${Region}:${Account}:document-classifier/${DocumentClassifierName}"},
		{Name: "comprehend_document_classifier_endpoint", Service: "comprehend", Resource: "document-classifier-endpoint", Template: "arn:${Partition}:comprehend:${Region}:${Account}:document-classifier-endpoint/${DocumentClassifierEndpointName}"},
		{Name: "comprehend_dominant_language_detection_job", Service: "comprehend", Resource: "dominant-language-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:dominant-language-detection-job/${JobId}"},
		{Name: "comprehend_entities_detection_job", Service: "comprehend", Resource: "entities-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:entities-detection-job/${JobId}"},
		{Name: "comprehend_entity_recognizer", Service: "comprehend", Resource: "entity-recognizer", Template: "arn:${Partition}:comprehend:${Region}:${Account}:entity-recognizer/${EntityRecognizerName}"},
		{Name: "comprehend_entity_recognizer_endpoint", Service: "comprehend", Resource: "entity-recognizer-endpoint", Template: "arn:${Partition}:comprehend:${Region}:${Account}:entity-recognizer-endpoint/${EntityRecognizerEndpointName}"},
		{Name: "comprehend_events_detection_job", Service: "comprehend", Resource: "events-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:events-detection-job/${JobId}"},
		{Name: "comprehend_flywheel", Service: "comprehend", Resource: "flywheel", Template: "arn:${Partition}:comprehend:${Region}:${Account}:flywheel/${FlywheelName}"},
		{Name: "comprehend_flywheel_dataset", Service: "comprehend", Resource: "flywheel-dataset", Template: "arn:${Partition}:comprehend:${Region}:${Account}:flywheel/${FlywheelName}/dataset/${DatasetName}"},
		{Name: "comprehend_key_phrases_detection_job", Service: "comprehend", Resource: "key-phrases-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:key-phrases-detection-job/${JobId}"},
		{Name: "comprehend_pii_entities_detection_job", Service: "comprehend", Resource: "pii-entities-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:pii-entities-detection-job/${JobId}"},
		{Name: "comprehend_sentiment_detection_job", Service: "comprehend", Resource: "sentiment-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:sentiment-detection-job/${JobId}"},
		{Name: "comprehend_targeted_sentiment_detection_job", Service: "comprehend", Resource: "targeted-sentiment-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:targeted-sentiment-detection-job/${JobId}"},
		{Name: "comprehend_topics_detection_job", Service: "comprehend", Resource: "topics-detection-job", Template: "arn:${Partition}:comprehend:${Region}:${Account}:topics-detection-job/${JobId}"},
	})
}
