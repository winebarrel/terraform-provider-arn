# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/resource-snapshot-job/identifier
output "partnercentral_resource_snapshot_job" {
  value = provider::arn::partnercentral_resource_snapshot_job("catalog", "identifier")
}
