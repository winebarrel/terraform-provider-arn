# arn:aws:ses:ap-northeast-1:111111111111:template/template-name
output "ses_template" {
  value = provider::arn::ses_template("template-name")
}
