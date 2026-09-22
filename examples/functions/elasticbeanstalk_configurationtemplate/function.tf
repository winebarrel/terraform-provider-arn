# arn:aws:elasticbeanstalk:ap-northeast-1:111111111111:configurationtemplate/application-name/template-name
output "elasticbeanstalk_configurationtemplate" {
  value = provider::arn::elasticbeanstalk_configurationtemplate("application-name", "template-name")
}
