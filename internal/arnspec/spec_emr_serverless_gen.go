// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: emr-serverless
// Source: https://servicereference.us-east-1.amazonaws.com/v1/emr-serverless/emr-serverless.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "emr_serverless_application", Service: "emr-serverless", Resource: "application", Template: "arn:${Partition}:emr-serverless:${Region}:${Account}:/applications/${ApplicationId}"},
		{Name: "emr_serverless_job_run", Service: "emr-serverless", Resource: "jobRun", Template: "arn:${Partition}:emr-serverless:${Region}:${Account}:/applications/${ApplicationId}/jobruns/${JobRunId}"},
		{Name: "emr_serverless_session", Service: "emr-serverless", Resource: "session", Template: "arn:${Partition}:emr-serverless:${Region}:${Account}:/applications/${ApplicationId}/sessions/${SessionId}"},
	})
}
