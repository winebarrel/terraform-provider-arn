# arn:aws:healthlake:ap-northeast-1:111111111111:datastore/fhir/datastore-id
output "healthlake_datastore" {
  value = provider::arn::healthlake_datastore("datastore-id")
}
