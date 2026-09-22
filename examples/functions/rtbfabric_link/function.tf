# arn:aws:rtbfabric:ap-northeast-1:111111111111:gateway/gateway-id/link/link-id
output "rtbfabric_link" {
  value = provider::arn::rtbfabric_link("gateway-id", "link-id")
}
