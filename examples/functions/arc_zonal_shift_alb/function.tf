# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:loadbalancer/app/load-balancer-name/load-balancer-id
output "arc_zonal_shift_alb" {
  value = provider::arn::arc_zonal_shift_alb("load-balancer-name", "load-balancer-id")
}
