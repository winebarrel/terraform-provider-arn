# arn:aws:vpc-lattice:ap-northeast-1:111111111111:service/service-id/listener/listener-id
output "vpc_lattice_listener" {
  value = provider::arn::vpc_lattice_listener("service-id", "listener-id")
}
