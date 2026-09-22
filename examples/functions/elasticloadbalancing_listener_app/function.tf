# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:listener/app/load-balancer-name/load-balancer-id/listener-id
output "elasticloadbalancing_listener_app" {
  value = provider::arn::elasticloadbalancing_listener_app("load-balancer-name", "load-balancer-id", "listener-id")
}
