// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cleanrooms-ml
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cleanrooms-ml/cleanrooms-ml.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cleanrooms_ml_audiencegenerationjob", Service: "cleanrooms-ml", Resource: "audiencegenerationjob", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:audience-generation-job/${ResourceId}"},
		{Name: "cleanrooms_ml_audiencemodel", Service: "cleanrooms-ml", Resource: "audiencemodel", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:audience-model/${ResourceId}"},
		{Name: "cleanrooms_ml_configured_model_algorithm", Service: "cleanrooms-ml", Resource: "ConfiguredModelAlgorithm", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:configured-model-algorithm/${ResourceId}"},
		{Name: "cleanrooms_ml_configured_model_algorithm_association", Service: "cleanrooms-ml", Resource: "ConfiguredModelAlgorithmAssociation", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:membership/${MembershipId}/configured-model-algorithm-association/${ResourceId}"},
		{Name: "cleanrooms_ml_configuredaudiencemodel", Service: "cleanrooms-ml", Resource: "configuredaudiencemodel", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:configured-audience-model/${ResourceId}"},
		{Name: "cleanrooms_ml_ml_input_channel", Service: "cleanrooms-ml", Resource: "MLInputChannel", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:membership/${MembershipId}/ml-input-channel/${ResourceId}"},
		{Name: "cleanrooms_ml_trained_model", Service: "cleanrooms-ml", Resource: "TrainedModel", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:membership/${MembershipId}/trained-model/${ResourceId}"},
		{Name: "cleanrooms_ml_trained_model_inference_job", Service: "cleanrooms-ml", Resource: "TrainedModelInferenceJob", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:membership/${MembershipId}/trained-model-inference-job/${ResourceId}"},
		{Name: "cleanrooms_ml_trainingdataset", Service: "cleanrooms-ml", Resource: "trainingdataset", Template: "arn:${Partition}:cleanrooms-ml:${Region}:${Account}:training-dataset/${ResourceId}"},
	})
}
