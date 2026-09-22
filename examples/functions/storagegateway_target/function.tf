# arn:aws:storagegateway:ap-northeast-1:111111111111:gateway/gateway-id/target/iscsi-target
output "storagegateway_target" {
  value = provider::arn::storagegateway_target("gateway-id", "iscsi-target")
}
