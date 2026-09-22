# arn:aws:vpc-lattice:ap-northeast-1:111111111111:service/service-id
output "vpc_lattice_svcs_tcp_service" {
  value = provider::arn::vpc_lattice_svcs_tcp_service("service-id")
}
