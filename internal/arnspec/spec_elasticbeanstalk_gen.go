// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elasticbeanstalk
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elasticbeanstalk/elasticbeanstalk.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elasticbeanstalk_application", Service: "elasticbeanstalk", Resource: "application", Template: "arn:${Partition}:elasticbeanstalk:${Region}:${Account}:application/${ApplicationName}"},
		{Name: "elasticbeanstalk_applicationversion", Service: "elasticbeanstalk", Resource: "applicationversion", Template: "arn:${Partition}:elasticbeanstalk:${Region}:${Account}:applicationversion/${ApplicationName}/${VersionLabel}"},
		{Name: "elasticbeanstalk_configurationtemplate", Service: "elasticbeanstalk", Resource: "configurationtemplate", Template: "arn:${Partition}:elasticbeanstalk:${Region}:${Account}:configurationtemplate/${ApplicationName}/${TemplateName}"},
		{Name: "elasticbeanstalk_environment", Service: "elasticbeanstalk", Resource: "environment", Template: "arn:${Partition}:elasticbeanstalk:${Region}:${Account}:environment/${ApplicationName}/${EnvironmentName}"},
		{Name: "elasticbeanstalk_platform", Service: "elasticbeanstalk", Resource: "platform", Template: "arn:${Partition}:elasticbeanstalk:${Region}::platform/${PlatformNameWithVersion}"},
		{Name: "elasticbeanstalk_solutionstack", Service: "elasticbeanstalk", Resource: "solutionstack", Template: "arn:${Partition}:elasticbeanstalk:${Region}::solutionstack/${SolutionStackName}"},
	})
}
