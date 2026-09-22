# arn:aws:partnercentral:ap-northeast-1:111111111111:catalog/catalog/engagement/engagement-identifier/resource/resource-type/resource-identifier/template/template-identifier/resource-snapshot/snapshot-revision
output "partnercentral_resource_snapshot" {
  value = provider::arn::partnercentral_resource_snapshot("catalog", "engagement-identifier", "resource-type", "resource-identifier", "template-identifier", "snapshot-revision")
}
