# arn:aws:rtbfabric:ap-northeast-1:111111111111:gateway/gateway-id
output "rtbfabric_responder_gateway" {
  value = provider::arn::rtbfabric_responder_gateway("gateway-id")
}
