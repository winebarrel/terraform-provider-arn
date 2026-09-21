// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: vpc-lattice-svcs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/vpc-lattice-svcs/vpc-lattice-svcs.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "vpc_lattice_svcs_service", Service: "vpc-lattice-svcs", Resource: "Service", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:service/${ServiceId}/${RequestPath}"},
		{Name: "vpc_lattice_svcs_tcp_service", Service: "vpc-lattice-svcs", Resource: "TCP Service", Template: "arn:${Partition}:vpc-lattice:${Region}:${Account}:service/${ServiceId}"},
	})
}
