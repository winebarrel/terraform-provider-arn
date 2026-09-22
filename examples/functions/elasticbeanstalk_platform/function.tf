# arn:aws:elasticbeanstalk:ap-northeast-1::platform/platform-name-with-version
output "elasticbeanstalk_platform" {
  value = provider::arn::elasticbeanstalk_platform("platform-name-with-version")
}
