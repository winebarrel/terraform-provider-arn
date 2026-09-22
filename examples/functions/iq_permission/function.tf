# arn:aws:iq-permission:ap-northeast-1::permission/permission-request-id
output "iq_permission" {
  value = provider::arn::iq_permission("permission-request-id")
}
