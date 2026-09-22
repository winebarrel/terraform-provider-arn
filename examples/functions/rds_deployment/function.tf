# arn:aws:rds:ap-northeast-1:111111111111:deployment:blue-green-deployment-identifier
output "rds_deployment" {
  value = provider::arn::rds_deployment("blue-green-deployment-identifier")
}
