# arn:aws:forecast:ap-northeast-1:111111111111:predictor-backtest-export-job/resource-id
output "forecast_predictor_backtest_export_job" {
  value = provider::arn::forecast_predictor_backtest_export_job("resource-id")
}
