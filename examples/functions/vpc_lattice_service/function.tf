# arn:aws:vpc-lattice:ap-northeast-1:111111111111:service/service-id
output "vpc_lattice_service" {
  value = provider::arn::vpc_lattice_service("service-id")
}
