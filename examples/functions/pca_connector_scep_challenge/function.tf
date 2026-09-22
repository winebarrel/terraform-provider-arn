# arn:aws:pca-connector-scep:ap-northeast-1:111111111111:connector/connector-id/challenge/challenge-id
output "pca_connector_scep_challenge" {
  value = provider::arn::pca_connector_scep_challenge("connector-id", "challenge-id")
}
