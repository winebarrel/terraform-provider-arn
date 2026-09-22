# arn:aws:ssm:ap-northeast-1:111111111111:cloud-connector/cloud-connector-id
output "ssm_cloud_connector" {
  value = provider::arn::ssm_cloud_connector("cloud-connector-id")
}
