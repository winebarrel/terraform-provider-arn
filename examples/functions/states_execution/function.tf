# arn:aws:states:ap-northeast-1:111111111111:execution:state-machine-name:execution-id
output "states_execution" {
  value = provider::arn::states_execution("state-machine-name", "execution-id")
}
