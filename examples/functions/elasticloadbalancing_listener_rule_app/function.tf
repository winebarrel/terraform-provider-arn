# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:listener-rule/app/load-balancer-name/load-balancer-id/listener-id/listener-rule-id
output "elasticloadbalancing_listener_rule_app" {
  value = provider::arn::elasticloadbalancing_listener_rule_app("load-balancer-name", "load-balancer-id", "listener-id", "listener-rule-id")
}
