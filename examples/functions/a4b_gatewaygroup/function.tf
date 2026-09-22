# arn:aws:a4b:ap-northeast-1:111111111111:gateway-group/resource-id
output "a4b_gatewaygroup" {
  value = provider::arn::a4b_gatewaygroup("resource-id")
}
