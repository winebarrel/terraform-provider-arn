// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: redshift
// Source: https://servicereference.us-east-1.amazonaws.com/v1/redshift/redshift.json
// Functions: 21
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "redshift_cluster", Service: "redshift", Resource: "cluster", Template: "arn:${Partition}:redshift:${Region}:${Account}:cluster:${ClusterName}"},
		{Name: "redshift_datashare", Service: "redshift", Resource: "datashare", Template: "arn:${Partition}:redshift:${Region}:${Account}:datashare:${ProducerClusterNamespace}/${DataShareName}"},
		{Name: "redshift_dbgroup", Service: "redshift", Resource: "dbgroup", Template: "arn:${Partition}:redshift:${Region}:${Account}:dbgroup:${ClusterName}/${DbGroup}"},
		{Name: "redshift_dbname", Service: "redshift", Resource: "dbname", Template: "arn:${Partition}:redshift:${Region}:${Account}:dbname:${ClusterName}/${DbName}"},
		{Name: "redshift_dbuser", Service: "redshift", Resource: "dbuser", Template: "arn:${Partition}:redshift:${Region}:${Account}:dbuser:${ClusterName}/${DbUser}"},
		{Name: "redshift_eventsubscription", Service: "redshift", Resource: "eventsubscription", Template: "arn:${Partition}:redshift:${Region}:${Account}:eventsubscription:${EventSubscriptionName}"},
		{Name: "redshift_hsmclientcertificate", Service: "redshift", Resource: "hsmclientcertificate", Template: "arn:${Partition}:redshift:${Region}:${Account}:hsmclientcertificate:${HSMClientCertificateId}"},
		{Name: "redshift_hsmconfiguration", Service: "redshift", Resource: "hsmconfiguration", Template: "arn:${Partition}:redshift:${Region}:${Account}:hsmconfiguration:${HSMConfigurationId}"},
		{Name: "redshift_integration", Service: "redshift", Resource: "integration", Template: "arn:${Partition}:redshift:${Region}:${Account}:integration:${IntegrationIdentifier}"},
		{Name: "redshift_namespace", Service: "redshift", Resource: "namespace", Template: "arn:${Partition}:redshift:${Region}:${Account}:namespace:${ClusterNamespace}"},
		{Name: "redshift_parametergroup", Service: "redshift", Resource: "parametergroup", Template: "arn:${Partition}:redshift:${Region}:${Account}:parametergroup:${ParameterGroupName}"},
		{Name: "redshift_qev2idcapplication", Service: "redshift", Resource: "qev2idcapplication", Template: "arn:${Partition}:redshift:${Region}:${Account}:qev2idcapplication:${Qev2IdcApplicationId}"},
		{Name: "redshift_redshiftidcapplication", Service: "redshift", Resource: "redshiftidcapplication", Template: "arn:${Partition}:redshift:${Region}:${Account}:redshiftidcapplication:${RedshiftIdcApplicationId}"},
		{Name: "redshift_securitygroup", Service: "redshift", Resource: "securitygroup", Template: "arn:${Partition}:redshift:${Region}:${Account}:securitygroup:${SecurityGroupName}/ec2securitygroup/${Owner}/${Ec2SecurityGroupId}"},
		{Name: "redshift_securitygroupingress_cidr", Service: "redshift", Resource: "securitygroupingress-cidr", Template: "arn:${Partition}:redshift:${Region}:${Account}:securitygroupingress:${SecurityGroupName}/cidrip/${IpRange}"},
		{Name: "redshift_securitygroupingress_ec2securitygroup", Service: "redshift", Resource: "securitygroupingress-ec2securitygroup", Template: "arn:${Partition}:redshift:${Region}:${Account}:securitygroupingress:${SecurityGroupName}/ec2securitygroup/${Owner}/${Ece2SecuritygroupId}"},
		{Name: "redshift_snapshot", Service: "redshift", Resource: "snapshot", Template: "arn:${Partition}:redshift:${Region}:${Account}:snapshot:${ClusterName}/${SnapshotName}"},
		{Name: "redshift_snapshotcopygrant", Service: "redshift", Resource: "snapshotcopygrant", Template: "arn:${Partition}:redshift:${Region}:${Account}:snapshotcopygrant:${SnapshotCopyGrantName}"},
		{Name: "redshift_snapshotschedule", Service: "redshift", Resource: "snapshotschedule", Template: "arn:${Partition}:redshift:${Region}:${Account}:snapshotschedule:${ScheduleIdentifier}"},
		{Name: "redshift_subnetgroup", Service: "redshift", Resource: "subnetgroup", Template: "arn:${Partition}:redshift:${Region}:${Account}:subnetgroup:${SubnetGroupName}"},
		{Name: "redshift_usagelimit", Service: "redshift", Resource: "usagelimit", Template: "arn:${Partition}:redshift:${Region}:${Account}:usagelimit:${UsageLimitId}"},
	})
}
