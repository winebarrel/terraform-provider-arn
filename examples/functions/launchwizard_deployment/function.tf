# arn:aws:launchwizard:ap-northeast-1:111111111111:deployment/deployment-id
output "launchwizard_deployment" {
  value = provider::arn::launchwizard_deployment("deployment-id")
}
