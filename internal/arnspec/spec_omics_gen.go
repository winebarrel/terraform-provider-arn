// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: omics
// Source: https://servicereference.us-east-1.amazonaws.com/v1/omics/omics.json
// Functions: 15
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "omics_annotation_store", Service: "omics", Resource: "AnnotationStore", Template: "arn:${Partition}:omics:${Region}:${Account}:annotationStore/${AnnotationStoreName}"},
		{Name: "omics_annotation_store_version", Service: "omics", Resource: "AnnotationStoreVersion", Template: "arn:${Partition}:omics:${Region}:${Account}:annotationStore/${AnnotationStoreName}/version/${AnnotationStoreVersionName}"},
		{Name: "omics_configuration", Service: "omics", Resource: "configuration", Template: "arn:${Partition}:omics:${Region}:${Account}:configuration/${Name}"},
		{Name: "omics_read_set", Service: "omics", Resource: "readSet", Template: "arn:${Partition}:omics:${Region}:${Account}:sequenceStore/${SequenceStoreId}/readSet/${ReadSetId}"},
		{Name: "omics_reference", Service: "omics", Resource: "reference", Template: "arn:${Partition}:omics:${Region}:${Account}:referenceStore/${ReferenceStoreId}/reference/${ReferenceId}"},
		{Name: "omics_reference_store", Service: "omics", Resource: "referenceStore", Template: "arn:${Partition}:omics:${Region}:${Account}:referenceStore/${ReferenceStoreId}"},
		{Name: "omics_run", Service: "omics", Resource: "run", Template: "arn:${Partition}:omics:${Region}:${Account}:run/${Id}"},
		{Name: "omics_run_batch", Service: "omics", Resource: "runBatch", Template: "arn:${Partition}:omics:${Region}:${Account}:runBatch/${BatchId}"},
		{Name: "omics_run_cache", Service: "omics", Resource: "runCache", Template: "arn:${Partition}:omics:${Region}:${Account}:runCache/${Id}"},
		{Name: "omics_run_group", Service: "omics", Resource: "runGroup", Template: "arn:${Partition}:omics:${Region}:${Account}:runGroup/${Id}"},
		{Name: "omics_sequence_store", Service: "omics", Resource: "sequenceStore", Template: "arn:${Partition}:omics:${Region}:${Account}:sequenceStore/${SequenceStoreId}"},
		{Name: "omics_task_resource", Service: "omics", Resource: "TaskResource", Template: "arn:${Partition}:omics:${Region}:${Account}:task/${Id}"},
		{Name: "omics_variant_store", Service: "omics", Resource: "VariantStore", Template: "arn:${Partition}:omics:${Region}:${Account}:variantStore/${VariantStoreName}"},
		{Name: "omics_workflow", Service: "omics", Resource: "workflow", Template: "arn:${Partition}:omics:${Region}:${Account}:workflow/${Id}"},
		{Name: "omics_workflow_version", Service: "omics", Resource: "WorkflowVersion", Template: "arn:${Partition}:omics:${Region}:${Account}:workflow/${Id}/version/${VersionName}"},
	})
}
