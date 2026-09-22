# arn:aws:medical-imaging:ap-northeast-1:111111111111:datastore/datastore-id/imageset/image-set-id
output "medical_imaging_imageset" {
  value = provider::arn::medical_imaging_imageset("datastore-id", "image-set-id")
}
