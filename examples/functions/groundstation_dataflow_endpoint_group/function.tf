# arn:aws:groundstation:ap-northeast-1:111111111111:dataflow-endpoint-group/dataflow-endpoint-group-id
output "groundstation_dataflow_endpoint_group" {
  value = provider::arn::groundstation_dataflow_endpoint_group("dataflow-endpoint-group-id")
}
