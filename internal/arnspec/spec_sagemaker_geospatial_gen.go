// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sagemaker-geospatial
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sagemaker-geospatial/sagemaker-geospatial.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sagemaker_geospatial_earth_observation_job", Service: "sagemaker-geospatial", Resource: "EarthObservationJob", Template: "arn:${Partition}:sagemaker-geospatial:${Region}:${Account}:earth-observation-job/${JobID}"},
		{Name: "sagemaker_geospatial_raster_data_collection", Service: "sagemaker-geospatial", Resource: "RasterDataCollection", Template: "arn:${Partition}:sagemaker-geospatial:${Region}:${Account}:raster-data-collection/${CollectionID}"},
		{Name: "sagemaker_geospatial_vector_enrichment_job", Service: "sagemaker-geospatial", Resource: "VectorEnrichmentJob", Template: "arn:${Partition}:sagemaker-geospatial:${Region}:${Account}:vector-enrichment-job/${JobID}"},
	})
}
