# arn:aws:medical-imaging:ap-northeast-1:111111111111:datastore/datastore-id
output "medical_imaging_datastore" {
  value = provider::arn::medical_imaging_datastore("datastore-id")
}
