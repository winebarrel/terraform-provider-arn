// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: lightsail
// Source: https://servicereference.us-east-1.amazonaws.com/v1/lightsail/lightsail.json
// Functions: 19
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "lightsail_alarm", Service: "lightsail", Resource: "Alarm", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Alarm/${Id}"},
		{Name: "lightsail_bucket", Service: "lightsail", Resource: "Bucket", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Bucket/${Id}"},
		{Name: "lightsail_certificate", Service: "lightsail", Resource: "Certificate", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Certificate/${Id}"},
		{Name: "lightsail_cloud_formation_stack_record", Service: "lightsail", Resource: "CloudFormationStackRecord", Template: "arn:${Partition}:lightsail:${Region}:${Account}:CloudFormationStackRecord/${Id}"},
		{Name: "lightsail_contact_method", Service: "lightsail", Resource: "ContactMethod", Template: "arn:${Partition}:lightsail:${Region}:${Account}:ContactMethod/${Id}"},
		{Name: "lightsail_container_service", Service: "lightsail", Resource: "ContainerService", Template: "arn:${Partition}:lightsail:${Region}:${Account}:ContainerService/${Id}"},
		{Name: "lightsail_disk", Service: "lightsail", Resource: "Disk", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Disk/${Id}"},
		{Name: "lightsail_disk_snapshot", Service: "lightsail", Resource: "DiskSnapshot", Template: "arn:${Partition}:lightsail:${Region}:${Account}:DiskSnapshot/${Id}"},
		{Name: "lightsail_distribution", Service: "lightsail", Resource: "Distribution", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Distribution/${Id}"},
		{Name: "lightsail_domain", Service: "lightsail", Resource: "Domain", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Domain/${Id}"},
		{Name: "lightsail_export_snapshot_record", Service: "lightsail", Resource: "ExportSnapshotRecord", Template: "arn:${Partition}:lightsail:${Region}:${Account}:ExportSnapshotRecord/${Id}"},
		{Name: "lightsail_instance", Service: "lightsail", Resource: "Instance", Template: "arn:${Partition}:lightsail:${Region}:${Account}:Instance/${Id}"},
		{Name: "lightsail_instance_snapshot", Service: "lightsail", Resource: "InstanceSnapshot", Template: "arn:${Partition}:lightsail:${Region}:${Account}:InstanceSnapshot/${Id}"},
		{Name: "lightsail_key_pair", Service: "lightsail", Resource: "KeyPair", Template: "arn:${Partition}:lightsail:${Region}:${Account}:KeyPair/${Id}"},
		{Name: "lightsail_load_balancer", Service: "lightsail", Resource: "LoadBalancer", Template: "arn:${Partition}:lightsail:${Region}:${Account}:LoadBalancer/${Id}"},
		{Name: "lightsail_load_balancer_tls_certificate", Service: "lightsail", Resource: "LoadBalancerTlsCertificate", Template: "arn:${Partition}:lightsail:${Region}:${Account}:LoadBalancerTlsCertificate/${Id}"},
		{Name: "lightsail_relational_database", Service: "lightsail", Resource: "RelationalDatabase", Template: "arn:${Partition}:lightsail:${Region}:${Account}:RelationalDatabase/${Id}"},
		{Name: "lightsail_relational_database_snapshot", Service: "lightsail", Resource: "RelationalDatabaseSnapshot", Template: "arn:${Partition}:lightsail:${Region}:${Account}:RelationalDatabaseSnapshot/${Id}"},
		{Name: "lightsail_static_ip", Service: "lightsail", Resource: "StaticIp", Template: "arn:${Partition}:lightsail:${Region}:${Account}:StaticIp/${Id}"},
	})
}
