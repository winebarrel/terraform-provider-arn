# arn:aws:rtbfabric:ap-northeast-1:111111111111:gateway/gateway-id/link/link-id/routing-rule/rule-id
output "rtbfabric_link_routing_rule" {
  value = provider::arn::rtbfabric_link_routing_rule("gateway-id", "link-id", "rule-id")
}
