// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: athena
// Source: https://servicereference.us-east-1.amazonaws.com/v1/athena/athena.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "athena_capacity_reservation", Service: "athena", Resource: "capacity-reservation", Template: "arn:${Partition}:athena:${Region}:${Account}:capacity-reservation/${CapacityReservationName}"},
		{Name: "athena_datacatalog", Service: "athena", Resource: "datacatalog", Template: "arn:${Partition}:athena:${Region}:${Account}:datacatalog/${DataCatalogName}"},
		{Name: "athena_session", Service: "athena", Resource: "session", Template: "arn:${Partition}:athena:${Region}:${Account}:workgroup/${WorkGroupName}/session/${SessionId}"},
		{Name: "athena_workgroup", Service: "athena", Resource: "workgroup", Template: "arn:${Partition}:athena:${Region}:${Account}:workgroup/${WorkGroupName}"},
	})
}
