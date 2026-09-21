// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: qapps
// Source: https://servicereference.us-east-1.amazonaws.com/v1/qapps/qapps.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "qapps_application", Service: "qapps", Resource: "application", Template: "arn:${Partition}:qbusiness:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "qapps_qapp", Service: "qapps", Resource: "qapp", Template: "arn:${Partition}:qapps:${Region}:${Account}:application/${ApplicationId}/qapp/${AppId}"},
		{Name: "qapps_qapp_session", Service: "qapps", Resource: "qapp-session", Template: "arn:${Partition}:qapps:${Region}:${Account}:application/${ApplicationId}/qapp/${AppId}/session/${SessionId}"},
	})
}
