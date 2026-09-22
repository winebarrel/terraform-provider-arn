# arn:aws:appflow:ap-northeast-1:111111111111:flow/flow-name
output "appflow_flow" {
  value = provider::arn::appflow_flow("flow-name")
}
