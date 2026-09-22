# arn:aws:elasticbeanstalk:ap-northeast-1::solutionstack/solution-stack-name
output "elasticbeanstalk_solutionstack" {
  value = provider::arn::elasticbeanstalk_solutionstack("solution-stack-name")
}
