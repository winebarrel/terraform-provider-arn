# arn:aws:imagebuilder:ap-northeast-1:111111111111:distribution-configuration/distribution-configuration-name
output "imagebuilder_distribution_configuration" {
  value = provider::arn::imagebuilder_distribution_configuration("distribution-configuration-name")
}
