// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: auditmanager
// Source: https://servicereference.us-east-1.amazonaws.com/v1/auditmanager/auditmanager.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "auditmanager_assessment", Service: "auditmanager", Resource: "assessment", Template: "arn:${Partition}:auditmanager:${Region}:${Account}:assessment/${AssessmentId}"},
		{Name: "auditmanager_assessment_control_set", Service: "auditmanager", Resource: "assessmentControlSet", Template: "arn:${Partition}:auditmanager:${Region}:${Account}:assessment/${AssessmentId}/controlSet/${ControlSetId}"},
		{Name: "auditmanager_assessment_framework", Service: "auditmanager", Resource: "assessmentFramework", Template: "arn:${Partition}:auditmanager:${Region}:${Account}:assessmentFramework/${AssessmentFrameworkId}"},
		{Name: "auditmanager_control", Service: "auditmanager", Resource: "control", Template: "arn:${Partition}:auditmanager:${Region}:${Account}:control/${ControlId}"},
	})
}
