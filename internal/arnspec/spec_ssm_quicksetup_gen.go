// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ssm-quicksetup
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ssm-quicksetup/ssm-quicksetup.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ssm_quicksetup_configuration_manager", Service: "ssm-quicksetup", Resource: "configuration-manager", Template: "arn:${Partition}:ssm-quicksetup:${Region}:${Account}:configuration-manager/${ConfigurationManagerId}"},
	})
}
