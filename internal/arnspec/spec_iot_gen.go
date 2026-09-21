// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iot
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iot/iot.json
// Functions: 33
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iot_authorizer", Service: "iot", Resource: "authorizer", Template: "arn:${Partition}:iot:${Region}:${Account}:authorizer/${AuthorizerName}"},
		{Name: "iot_billinggroup", Service: "iot", Resource: "billinggroup", Template: "arn:${Partition}:iot:${Region}:${Account}:billinggroup/${BillingGroupName}"},
		{Name: "iot_cacert", Service: "iot", Resource: "cacert", Template: "arn:${Partition}:iot:${Region}:${Account}:cacert/${CACertificate}"},
		{Name: "iot_cert", Service: "iot", Resource: "cert", Template: "arn:${Partition}:iot:${Region}:${Account}:cert/${Certificate}"},
		{Name: "iot_certificateprovider", Service: "iot", Resource: "certificateprovider", Template: "arn:${Partition}:iot:${Region}:${Account}:certificateprovider/${CertificateProviderName}"},
		{Name: "iot_client", Service: "iot", Resource: "client", Template: "arn:${Partition}:iot:${Region}:${Account}:client/${ClientId}"},
		{Name: "iot_command", Service: "iot", Resource: "command", Template: "arn:${Partition}:iot:${Region}:${Account}:command/${CommandId}"},
		{Name: "iot_custommetric", Service: "iot", Resource: "custommetric", Template: "arn:${Partition}:iot:${Region}:${Account}:custommetric/${MetricName}"},
		{Name: "iot_destination", Service: "iot", Resource: "destination", Template: "arn:${Partition}:iot:${Region}:${Account}:ruledestination/${DestinationType}/${Uuid}"},
		{Name: "iot_dimension", Service: "iot", Resource: "dimension", Template: "arn:${Partition}:iot:${Region}:${Account}:dimension/${DimensionName}"},
		{Name: "iot_domainconfiguration", Service: "iot", Resource: "domainconfiguration", Template: "arn:${Partition}:iot:${Region}:${Account}:domainconfiguration/${DomainConfigurationName}/${Id}"},
		{Name: "iot_dynamicthinggroup", Service: "iot", Resource: "dynamicthinggroup", Template: "arn:${Partition}:iot:${Region}:${Account}:thinggroup/${ThingGroupName}"},
		{Name: "iot_fleetmetric", Service: "iot", Resource: "fleetmetric", Template: "arn:${Partition}:iot:${Region}:${Account}:fleetmetric/${FleetMetricName}"},
		{Name: "iot_index", Service: "iot", Resource: "index", Template: "arn:${Partition}:iot:${Region}:${Account}:index/${IndexName}"},
		{Name: "iot_job", Service: "iot", Resource: "job", Template: "arn:${Partition}:iot:${Region}:${Account}:job/${JobId}"},
		{Name: "iot_jobtemplate", Service: "iot", Resource: "jobtemplate", Template: "arn:${Partition}:iot:${Region}:${Account}:jobtemplate/${JobTemplateId}"},
		{Name: "iot_mitigationaction", Service: "iot", Resource: "mitigationaction", Template: "arn:${Partition}:iot:${Region}:${Account}:mitigationaction/${MitigationActionName}"},
		{Name: "iot_otaupdate", Service: "iot", Resource: "otaupdate", Template: "arn:${Partition}:iot:${Region}:${Account}:otaupdate/${OtaUpdateId}"},
		{Name: "iot_package", Service: "iot", Resource: "package", Template: "arn:${Partition}:iot:${Region}:${Account}:package/${PackageName}"},
		{Name: "iot_packageversion", Service: "iot", Resource: "packageversion", Template: "arn:${Partition}:iot:${Region}:${Account}:package/${PackageName}/version/${VersionName}"},
		{Name: "iot_policy", Service: "iot", Resource: "policy", Template: "arn:${Partition}:iot:${Region}:${Account}:policy/${PolicyName}"},
		{Name: "iot_provisioningtemplate", Service: "iot", Resource: "provisioningtemplate", Template: "arn:${Partition}:iot:${Region}:${Account}:provisioningtemplate/${ProvisioningTemplate}"},
		{Name: "iot_rolealias", Service: "iot", Resource: "rolealias", Template: "arn:${Partition}:iot:${Region}:${Account}:rolealias/${RoleAlias}"},
		{Name: "iot_rule", Service: "iot", Resource: "rule", Template: "arn:${Partition}:iot:${Region}:${Account}:rule/${RuleName}"},
		{Name: "iot_scheduledaudit", Service: "iot", Resource: "scheduledaudit", Template: "arn:${Partition}:iot:${Region}:${Account}:scheduledaudit/${ScheduleName}"},
		{Name: "iot_securityprofile", Service: "iot", Resource: "securityprofile", Template: "arn:${Partition}:iot:${Region}:${Account}:securityprofile/${SecurityProfileName}"},
		{Name: "iot_stream", Service: "iot", Resource: "stream", Template: "arn:${Partition}:iot:${Region}:${Account}:stream/${StreamId}"},
		{Name: "iot_thing", Service: "iot", Resource: "thing", Template: "arn:${Partition}:iot:${Region}:${Account}:thing/${ThingName}"},
		{Name: "iot_thinggroup", Service: "iot", Resource: "thinggroup", Template: "arn:${Partition}:iot:${Region}:${Account}:thinggroup/${ThingGroupName}"},
		{Name: "iot_thingtype", Service: "iot", Resource: "thingtype", Template: "arn:${Partition}:iot:${Region}:${Account}:thingtype/${ThingTypeName}"},
		{Name: "iot_topic", Service: "iot", Resource: "topic", Template: "arn:${Partition}:iot:${Region}:${Account}:topic/${TopicName}"},
		{Name: "iot_topicfilter", Service: "iot", Resource: "topicfilter", Template: "arn:${Partition}:iot:${Region}:${Account}:topicfilter/${TopicFilter}"},
		{Name: "iot_tunnel", Service: "iot", Resource: "tunnel", Template: "arn:${Partition}:iot:${Region}:${Account}:tunnel/${TunnelId}"},
	})
}
