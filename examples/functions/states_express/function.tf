# arn:aws:states:ap-northeast-1:111111111111:express:state-machine-name:execution-id:express-id
output "states_express" {
  value = provider::arn::states_express("state-machine-name", "execution-id", "express-id")
}
