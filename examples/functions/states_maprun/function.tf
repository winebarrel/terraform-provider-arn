# arn:aws:states:ap-northeast-1:111111111111:mapRun:state-machine-name/map-run-label:map-run-id
output "states_maprun" {
  value = provider::arn::states_maprun("state-machine-name", "map-run-label", "map-run-id")
}
