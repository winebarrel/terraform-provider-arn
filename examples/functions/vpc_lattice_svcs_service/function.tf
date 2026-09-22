# arn:aws:vpc-lattice:ap-northeast-1:111111111111:service/service-id/request-path
output "vpc_lattice_svcs_service" {
  value = provider::arn::vpc_lattice_svcs_service("service-id", "request-path")
}
