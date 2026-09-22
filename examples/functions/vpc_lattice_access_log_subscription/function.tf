# arn:aws:vpc-lattice:ap-northeast-1:111111111111:accesslogsubscription/access-log-subscription-id
output "vpc_lattice_access_log_subscription" {
  value = provider::arn::vpc_lattice_access_log_subscription("access-log-subscription-id")
}
