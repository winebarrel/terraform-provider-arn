# arn:aws:vpc-lattice:ap-northeast-1:111111111111:domainverification/domain-verification-id
output "vpc_lattice_domain_verification" {
  value = provider::arn::vpc_lattice_domain_verification("domain-verification-id")
}
