# arn:aws:vpc-lattice:ap-northeast-1:111111111111:service/service-id/listener/listener-id/rule/rule-id
output "vpc_lattice_rule" {
  value = provider::arn::vpc_lattice_rule("service-id", "listener-id", "rule-id")
}
