# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:listener/gwy/load-balancer-name/load-balancer-id/listener-id
output "elasticloadbalancing_listener_gwy" {
  value = provider::arn::elasticloadbalancing_listener_gwy("load-balancer-name", "load-balancer-id", "listener-id")
}
