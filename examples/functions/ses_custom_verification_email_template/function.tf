# arn:aws:ses:ap-northeast-1:111111111111:custom-verification-email-template/template-name
output "ses_custom_verification_email_template" {
  value = provider::arn::ses_custom_verification_email_template("template-name")
}
