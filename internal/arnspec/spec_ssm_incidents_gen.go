// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ssm-incidents
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ssm-incidents/ssm-incidents.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ssm_incidents_incident_record", Service: "ssm-incidents", Resource: "incident-record", Template: "arn:${Partition}:ssm-incidents::${Account}:incident-record/${ResponsePlan}/${IncidentRecord}"},
		{Name: "ssm_incidents_replication_set", Service: "ssm-incidents", Resource: "replication-set", Template: "arn:${Partition}:ssm-incidents::${Account}:replication-set/${ReplicationSet}"},
		{Name: "ssm_incidents_response_plan", Service: "ssm-incidents", Resource: "response-plan", Template: "arn:${Partition}:ssm-incidents::${Account}:response-plan/${ResponsePlan}"},
	})
}
