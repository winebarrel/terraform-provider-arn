# arn:aws:entityresolution:ap-northeast-1:111111111111:idnamespace/id-namespace-name
output "entityresolution_id_namespace" {
  value = provider::arn::entityresolution_id_namespace("id-namespace-name")
}
