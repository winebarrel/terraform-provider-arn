# arn:aws:mediaconnect:ap-northeast-1:111111111111:entitlement:flow-id:entitlement-name
output "mediaconnect_entitlement" {
  value = provider::arn::mediaconnect_entitlement("flow-id", "entitlement-name")
}
