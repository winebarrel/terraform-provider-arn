// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: states
// Source: https://servicereference.us-east-1.amazonaws.com/v1/states/states.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "states_activity", Service: "states", Resource: "activity", Template: "arn:${Partition}:states:${Region}:${Account}:activity:${ActivityName}"},
		{Name: "states_execution", Service: "states", Resource: "execution", Template: "arn:${Partition}:states:${Region}:${Account}:execution:${StateMachineName}:${ExecutionId}"},
		{Name: "states_express", Service: "states", Resource: "express", Template: "arn:${Partition}:states:${Region}:${Account}:express:${StateMachineName}:${ExecutionId}:${ExpressId}"},
		{Name: "states_labelled_execution", Service: "states", Resource: "labelled execution", Template: "arn:${Partition}:states:${Region}:${Account}:execution:${StateMachineName}/${MapRunLabel}:${ExecutionId}"},
		{Name: "states_labelled_express", Service: "states", Resource: "labelled express", Template: "arn:${Partition}:states:${Region}:${Account}:express:${StateMachineName}/${MapRunLabel}:${ExecutionId}:${ExpressId}"},
		{Name: "states_maprun", Service: "states", Resource: "maprun", Template: "arn:${Partition}:states:${Region}:${Account}:mapRun:${StateMachineName}/${MapRunLabel}:${MapRunId}"},
		{Name: "states_statemachine", Service: "states", Resource: "statemachine", Template: "arn:${Partition}:states:${Region}:${Account}:stateMachine:${StateMachineName}"},
		{Name: "states_statemachinealias", Service: "states", Resource: "statemachinealias", Template: "arn:${Partition}:states:${Region}:${Account}:stateMachine:${StateMachineName}:${StateMachineAliasName}"},
		{Name: "states_statemachineversion", Service: "states", Resource: "statemachineversion", Template: "arn:${Partition}:states:${Region}:${Account}:stateMachine:${StateMachineName}:${StateMachineVersionId}"},
	})
}
