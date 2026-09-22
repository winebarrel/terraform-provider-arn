# arn:aws:mobiletargeting:ap-northeast-1:111111111111:templates/template-name/template-type
output "mobiletargeting_template" {
  value = provider::arn::mobiletargeting_template("template-name", "template-type")
}
