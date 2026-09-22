# arn:aws:vpc-lattice:ap-northeast-1:111111111111:targetgroup/target-group-id
output "vpc_lattice_target_group" {
  value = provider::arn::vpc_lattice_target_group("target-group-id")
}
