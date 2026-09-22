# arn:aws:cloudformation:ap-northeast-1:111111111111:resourceScan/id
output "cloudformation_resourcescan" {
  value = provider::arn::cloudformation_resourcescan("id")
}
