// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: braket
// Source: https://servicereference.us-east-1.amazonaws.com/v1/braket/braket.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "braket_device", Service: "braket", Resource: "device", Template: "arn:${Partition}:braket:*:*:device/${DeviceType}/${Provider}/${DeviceId}"},
		{Name: "braket_job", Service: "braket", Resource: "job", Template: "arn:${Partition}:braket:${Region}:${Account}:job/${RandomId}"},
		{Name: "braket_quantum_task", Service: "braket", Resource: "quantum-task", Template: "arn:${Partition}:braket:${Region}:${Account}:quantum-task/${RandomId}"},
		{Name: "braket_spending_limit", Service: "braket", Resource: "spending-limit", Template: "arn:${Partition}:braket:${Region}:${Account}:spending-limit/${RandomId}"},
	})
}
