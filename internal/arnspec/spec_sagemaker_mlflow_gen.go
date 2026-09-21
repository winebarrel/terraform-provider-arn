// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sagemaker-mlflow
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sagemaker-mlflow/sagemaker-mlflow.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sagemaker_mlflow_mlflow_tracking_server", Service: "sagemaker-mlflow", Resource: "mlflow-tracking-server", Template: "arn:${Partition}:sagemaker:${Region}:${Account}:mlflow-tracking-server/${MlflowTrackingServerName}"},
	})
}
