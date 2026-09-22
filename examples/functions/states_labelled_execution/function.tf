# arn:aws:states:ap-northeast-1:111111111111:execution:state-machine-name/map-run-label:execution-id
output "states_labelled_execution" {
  value = provider::arn::states_labelled_execution("state-machine-name", "map-run-label", "execution-id")
}
