// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: airflow
// Source: https://servicereference.us-east-1.amazonaws.com/v1/airflow/airflow.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "airflow_environment", Service: "airflow", Resource: "environment", Template: "arn:${Partition}:airflow:${Region}:${Account}:environment/${EnvironmentName}"},
		{Name: "airflow_rbac_role", Service: "airflow", Resource: "rbac-role", Template: "arn:${Partition}:airflow:${Region}:${Account}:role/${EnvironmentName}/${RoleName}"},
	})
}
