# arn:aws:elasticbeanstalk:ap-northeast-1:111111111111:applicationversion/application-name/version-label
output "elasticbeanstalk_applicationversion" {
  value = provider::arn::elasticbeanstalk_applicationversion("application-name", "version-label")
}
