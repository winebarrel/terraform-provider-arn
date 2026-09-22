# arn:aws:odb:ap-northeast-1:111111111111:cloud-exadata-infrastructure/cloud-exadata-infrastructure-id
output "odb_cloud_exadata_infrastructure" {
  value = provider::arn::odb_cloud_exadata_infrastructure("cloud-exadata-infrastructure-id")
}
