# arn:aws:quicksight:ap-northeast-1:111111111111:email-customization-template/resource-id
output "quicksight_email_customization_template" {
  value = provider::arn::quicksight_email_customization_template("resource-id")
}
