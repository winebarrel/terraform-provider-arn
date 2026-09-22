# arn:aws:rtbfabric:ap-northeast-1:111111111111:gateway/gateway-id
output "rtbfabric_requester_gateway" {
  value = provider::arn::rtbfabric_requester_gateway("gateway-id")
}
