# arn:aws:states:ap-northeast-1:111111111111:stateMachine:state-machine-name:state-machine-alias-name
output "states_statemachinealias" {
  value = provider::arn::states_statemachinealias("state-machine-name", "state-machine-alias-name")
}
