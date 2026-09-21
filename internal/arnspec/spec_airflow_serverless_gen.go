// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: airflow-serverless
// Source: https://servicereference.us-east-1.amazonaws.com/v1/airflow-serverless/airflow-serverless.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "airflow_serverless_workflow", Service: "airflow-serverless", Resource: "Workflow", Template: "arn:${Partition}:airflow-serverless:${Region}:${Account}:workflow/${WorkflowId}"},
	})
}
