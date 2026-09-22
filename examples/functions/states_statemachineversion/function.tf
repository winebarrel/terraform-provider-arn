# arn:aws:states:ap-northeast-1:111111111111:stateMachine:state-machine-name:state-machine-version-id
output "states_statemachineversion" {
  value = provider::arn::states_statemachineversion("state-machine-name", "state-machine-version-id")
}
