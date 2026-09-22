# arn:aws:sagemaker-geospatial:ap-northeast-1:111111111111:raster-data-collection/collection-id
output "sagemaker_geospatial_raster_data_collection" {
  value = provider::arn::sagemaker_geospatial_raster_data_collection("collection-id")
}
