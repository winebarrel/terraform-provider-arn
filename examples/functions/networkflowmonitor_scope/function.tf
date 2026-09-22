# arn:aws:networkflowmonitor:ap-northeast-1:111111111111:scope/scope-id
output "networkflowmonitor_scope" {
  value = provider::arn::networkflowmonitor_scope("scope-id")
}
