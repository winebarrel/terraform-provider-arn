# arn:aws:quicksight:ap-northeast-1:111111111111:template/resource-id
output "quicksight_template" {
  value = provider::arn::quicksight_template("resource-id")
}
