# arn:aws:proton:ap-northeast-1:111111111111:environment-template/name
output "proton_environment_template" {
  value = provider::arn::proton_environment_template("name")
}
