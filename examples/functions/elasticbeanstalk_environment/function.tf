# arn:aws:elasticbeanstalk:ap-northeast-1:111111111111:environment/application-name/environment-name
output "elasticbeanstalk_environment" {
  value = provider::arn::elasticbeanstalk_environment("application-name", "environment-name")
}
