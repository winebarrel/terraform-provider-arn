# arn:aws:proton:ap-northeast-1:111111111111:deployment/id
output "proton_deployment" {
  value = provider::arn::proton_deployment("id")
}
