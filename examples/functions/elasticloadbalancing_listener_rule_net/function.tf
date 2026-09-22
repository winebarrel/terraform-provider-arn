# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:listener-rule/net/load-balancer-name/load-balancer-id/listener-id/listener-rule-id
output "elasticloadbalancing_listener_rule_net" {
  value = provider::arn::elasticloadbalancing_listener_rule_net("load-balancer-name", "load-balancer-id", "listener-id", "listener-rule-id")
}
