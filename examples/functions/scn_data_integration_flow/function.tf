# arn:aws:scn:ap-northeast-1:111111111111:instance/instance-id/data-integration-flows/flow-name
output "scn_data_integration_flow" {
  value = provider::arn::scn_data_integration_flow("instance-id", "flow-name")
}
