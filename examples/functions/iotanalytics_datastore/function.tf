# arn:aws:iotanalytics:ap-northeast-1:111111111111:datastore/datastore-name
output "iotanalytics_datastore" {
  value = provider::arn::iotanalytics_datastore("datastore-name")
}
