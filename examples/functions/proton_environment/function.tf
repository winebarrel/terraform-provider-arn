# arn:aws:proton:ap-northeast-1:111111111111:environment/name
output "proton_environment" {
  value = provider::arn::proton_environment("name")
}
