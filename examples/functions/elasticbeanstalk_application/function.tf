# arn:aws:elasticbeanstalk:ap-northeast-1:111111111111:application/application-name
output "elasticbeanstalk_application" {
  value = provider::arn::elasticbeanstalk_application("application-name")
}
