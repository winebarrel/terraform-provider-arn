# arn:aws:states:ap-northeast-1:111111111111:stateMachine:state-machine-name
output "states_statemachine" {
  value = provider::arn::states_statemachine("state-machine-name")
}
