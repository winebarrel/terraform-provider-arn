// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: machinelearning
// Source: https://servicereference.us-east-1.amazonaws.com/v1/machinelearning/machinelearning.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "machinelearning_batchprediction", Service: "machinelearning", Resource: "batchprediction", Template: "arn:${Partition}:machinelearning:${Region}:${Account}:batchprediction/${BatchPredictionId}"},
		{Name: "machinelearning_datasource", Service: "machinelearning", Resource: "datasource", Template: "arn:${Partition}:machinelearning:${Region}:${Account}:datasource/${DatasourceId}"},
		{Name: "machinelearning_evaluation", Service: "machinelearning", Resource: "evaluation", Template: "arn:${Partition}:machinelearning:${Region}:${Account}:evaluation/${EvaluationId}"},
		{Name: "machinelearning_mlmodel", Service: "machinelearning", Resource: "mlmodel", Template: "arn:${Partition}:machinelearning:${Region}:${Account}:mlmodel/${MlModelId}"},
	})
}
