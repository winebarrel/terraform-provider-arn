# arn:aws:redshift:ap-northeast-1:111111111111:qev2idcapplication:qev2-idc-application-id
output "redshift_qev2idcapplication" {
  value = provider::arn::redshift_qev2idcapplication("qev2-idc-application-id")
}
