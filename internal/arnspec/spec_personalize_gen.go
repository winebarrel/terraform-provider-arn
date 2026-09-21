// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: personalize
// Source: https://servicereference.us-east-1.amazonaws.com/v1/personalize/personalize.json
// Functions: 18
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "personalize_algorithm", Service: "personalize", Resource: "algorithm", Template: "arn:${Partition}:personalize:::algorithm/${ResourceId}"},
		{Name: "personalize_batch_inference_job", Service: "personalize", Resource: "batchInferenceJob", Template: "arn:${Partition}:personalize:${Region}:${Account}:batch-inference-job/${ResourceId}"},
		{Name: "personalize_batch_segment_job", Service: "personalize", Resource: "batchSegmentJob", Template: "arn:${Partition}:personalize:${Region}:${Account}:batch-segment-job/${ResourceId}"},
		{Name: "personalize_campaign", Service: "personalize", Resource: "campaign", Template: "arn:${Partition}:personalize:${Region}:${Account}:campaign/${ResourceId}"},
		{Name: "personalize_data_deletion_job", Service: "personalize", Resource: "dataDeletionJob", Template: "arn:${Partition}:personalize:${Region}:${Account}:data-deletion-job/${ResourceId}"},
		{Name: "personalize_data_insights_job", Service: "personalize", Resource: "dataInsightsJob", Template: "arn:${Partition}:personalize:${Region}:${Account}:data-insights-job/${ResourceId}"},
		{Name: "personalize_dataset", Service: "personalize", Resource: "dataset", Template: "arn:${Partition}:personalize:${Region}:${Account}:dataset/${ResourceId}"},
		{Name: "personalize_dataset_export_job", Service: "personalize", Resource: "datasetExportJob", Template: "arn:${Partition}:personalize:${Region}:${Account}:dataset-export-job/${ResourceId}"},
		{Name: "personalize_dataset_group", Service: "personalize", Resource: "datasetGroup", Template: "arn:${Partition}:personalize:${Region}:${Account}:dataset-group/${ResourceId}"},
		{Name: "personalize_dataset_import_job", Service: "personalize", Resource: "datasetImportJob", Template: "arn:${Partition}:personalize:${Region}:${Account}:dataset-import-job/${ResourceId}"},
		{Name: "personalize_event_tracker", Service: "personalize", Resource: "eventTracker", Template: "arn:${Partition}:personalize:${Region}:${Account}:event-tracker/${ResourceId}"},
		{Name: "personalize_feature_transformation", Service: "personalize", Resource: "featureTransformation", Template: "arn:${Partition}:personalize:::feature-transformation/${ResourceId}"},
		{Name: "personalize_filter", Service: "personalize", Resource: "filter", Template: "arn:${Partition}:personalize:${Region}:${Account}:filter/${ResourceId}"},
		{Name: "personalize_metric_attribution", Service: "personalize", Resource: "metricAttribution", Template: "arn:${Partition}:personalize:${Region}:${Account}:metric-attribution/${ResourceId}"},
		{Name: "personalize_recipe", Service: "personalize", Resource: "recipe", Template: "arn:${Partition}:personalize:::recipe/${ResourceId}"},
		{Name: "personalize_recommender", Service: "personalize", Resource: "recommender", Template: "arn:${Partition}:personalize:${Region}:${Account}:recommender/${ResourceId}"},
		{Name: "personalize_schema", Service: "personalize", Resource: "schema", Template: "arn:${Partition}:personalize:${Region}:${Account}:schema/${ResourceId}"},
		{Name: "personalize_solution", Service: "personalize", Resource: "solution", Template: "arn:${Partition}:personalize:${Region}:${Account}:solution/${ResourceId}"},
	})
}
