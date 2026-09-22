# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:targetgroup/target-group-name/target-group-id
output "elasticloadbalancing_targetgroup" {
  value = provider::arn::elasticloadbalancing_targetgroup("target-group-name", "target-group-id")
}
