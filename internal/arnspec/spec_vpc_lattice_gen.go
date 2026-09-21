// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: vpc-lattice
// Source: https://servicereference.us-east-1.amazonaws.com/v1/vpc-lattice/vpc-lattice.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "vpc_lattice_access_log_subscription", Service: "vpc-lattice", Resource: "AccessLogSubscription", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:accesslogsubscription/${AccessLogSubscriptionId}"},
		{Name: "vpc_lattice_domain_verification", Service: "vpc-lattice", Resource: "DomainVerification", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:domainverification/${DomainVerificationId}"},
		{Name: "vpc_lattice_listener", Service: "vpc-lattice", Resource: "Listener", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:service/${ServiceId}/listener/${ListenerId}"},
		{Name: "vpc_lattice_resource_configuration", Service: "vpc-lattice", Resource: "ResourceConfiguration", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:resourceconfiguration/${ResourceConfigurationId}"},
		{Name: "vpc_lattice_resource_endpoint_association", Service: "vpc-lattice", Resource: "ResourceEndpointAssociation", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:resourceendpointassociation/${ResourceEndpointAssociationId}"},
		{Name: "vpc_lattice_resource_gateway", Service: "vpc-lattice", Resource: "ResourceGateway", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:resourcegateway/${ResourceGatewayId}"},
		{Name: "vpc_lattice_rule", Service: "vpc-lattice", Resource: "Rule", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:service/${ServiceId}/listener/${ListenerId}/rule/${RuleId}"},
		{Name: "vpc_lattice_service", Service: "vpc-lattice", Resource: "Service", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:service/${ServiceId}"},
		{Name: "vpc_lattice_service_network", Service: "vpc-lattice", Resource: "ServiceNetwork", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:servicenetwork/${ServiceNetworkId}"},
		{Name: "vpc_lattice_service_network_resource_association", Service: "vpc-lattice", Resource: "ServiceNetworkResourceAssociation", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:servicenetworkresourceassociation/${ServiceNetworkResourceAssociationId}"},
		{Name: "vpc_lattice_service_network_service_association", Service: "vpc-lattice", Resource: "ServiceNetworkServiceAssociation", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:servicenetworkserviceassociation/${ServiceNetworkServiceAssociationId}"},
		{Name: "vpc_lattice_service_network_vpc_association", Service: "vpc-lattice", Resource: "ServiceNetworkVpcAssociation", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:servicenetworkvpcassociation/${ServiceNetworkVpcAssociationId}"},
		{Name: "vpc_lattice_target_group", Service: "vpc-lattice", Resource: "TargetGroup", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:targetgroup/${TargetGroupId}"},
	})
}
