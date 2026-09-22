# arn:aws:appsync:ap-northeast-1:111111111111:apis/graph-qlapi-id
output "wafv2_appsync" {
  value = provider::arn::wafv2_appsync("graph-qlapi-id")
}
