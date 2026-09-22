# arn:aws:lookoutvision:ap-northeast-1:111111111111:model/project-name/model-version
output "lookoutvision_model" {
  value = provider::arn::lookoutvision_model("project-name", "model-version")
}
