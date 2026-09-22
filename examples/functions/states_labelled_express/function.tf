# arn:aws:states:ap-northeast-1:111111111111:express:state-machine-name/map-run-label:execution-id:express-id
output "states_labelled_express" {
  value = provider::arn::states_labelled_express("state-machine-name", "map-run-label", "execution-id", "express-id")
}
