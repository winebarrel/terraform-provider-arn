# arn:aws:iotsitewise:ap-northeast-1:111111111111:time-series/time-series-id
output "iotsitewise_time_series" {
  value = provider::arn::iotsitewise_time_series("time-series-id")
}
