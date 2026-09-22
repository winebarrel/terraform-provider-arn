# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:listener/net/load-balancer-name/load-balancer-id/listener-id
output "elasticloadbalancing_listener_net" {
  value = provider::arn::elasticloadbalancing_listener_net("load-balancer-name", "load-balancer-id", "listener-id")
}
