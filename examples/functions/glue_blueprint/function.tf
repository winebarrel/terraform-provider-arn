# arn:aws:glue:ap-northeast-1:111111111111:blueprint/blueprint-name
output "glue_blueprint" {
  value = provider::arn::glue_blueprint("blueprint-name")
}
