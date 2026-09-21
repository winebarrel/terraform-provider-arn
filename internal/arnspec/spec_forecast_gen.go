// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: forecast
// Source: https://servicereference.us-east-1.amazonaws.com/v1/forecast/forecast.json
// Functions: 15
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "forecast_algorithm", Service: "forecast", Resource: "algorithm", Template: "arn:${Partition}:forecast:::algorithm/${ResourceId}"},
		{Name: "forecast_dataset", Service: "forecast", Resource: "dataset", Template: "arn:${Partition}:forecast:${Region}:${Account}:dataset/${ResourceId}"},
		{Name: "forecast_dataset_group", Service: "forecast", Resource: "datasetGroup", Template: "arn:${Partition}:forecast:${Region}:${Account}:dataset-group/${ResourceId}"},
		{Name: "forecast_dataset_import_job", Service: "forecast", Resource: "datasetImportJob", Template: "arn:${Partition}:forecast:${Region}:${Account}:dataset-import-job/${ResourceId}"},
		{Name: "forecast_endpoint", Service: "forecast", Resource: "endpoint", Template: "arn:${Partition}:forecast:${Region}:${Account}:forecast-endpoint/${ResourceId}"},
		{Name: "forecast_explainability", Service: "forecast", Resource: "explainability", Template: "arn:${Partition}:forecast:${Region}:${Account}:explainability/${ResourceId}"},
		{Name: "forecast_explainability_export", Service: "forecast", Resource: "explainabilityExport", Template: "arn:${Partition}:forecast:${Region}:${Account}:explainability-export/${ResourceId}"},
		{Name: "forecast_forecast", Service: "forecast", Resource: "forecast", Template: "arn:${Partition}:forecast:${Region}:${Account}:forecast/${ResourceId}"},
		{Name: "forecast_forecast_export", Service: "forecast", Resource: "forecastExport", Template: "arn:${Partition}:forecast:${Region}:${Account}:forecast-export-job/${ResourceId}"},
		{Name: "forecast_monitor", Service: "forecast", Resource: "monitor", Template: "arn:${Partition}:forecast:${Region}:${Account}:monitor/${ResourceId}"},
		{Name: "forecast_predictor", Service: "forecast", Resource: "predictor", Template: "arn:${Partition}:forecast:${Region}:${Account}:predictor/${ResourceId}"},
		{Name: "forecast_predictor_backtest_export_job", Service: "forecast", Resource: "predictorBacktestExportJob", Template: "arn:${Partition}:forecast:${Region}:${Account}:predictor-backtest-export-job/${ResourceId}"},
		{Name: "forecast_what_if_analysis", Service: "forecast", Resource: "whatIfAnalysis", Template: "arn:${Partition}:forecast:${Region}:${Account}:what-if-analysis/${ResourceId}"},
		{Name: "forecast_what_if_forecast", Service: "forecast", Resource: "whatIfForecast", Template: "arn:${Partition}:forecast:${Region}:${Account}:what-if-forecast/${ResourceId}"},
		{Name: "forecast_what_if_forecast_export", Service: "forecast", Resource: "whatIfForecastExport", Template: "arn:${Partition}:forecast:${Region}:${Account}:what-if-forecast-export/${ResourceId}"},
	})
}
