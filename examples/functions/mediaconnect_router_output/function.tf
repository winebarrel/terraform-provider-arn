# arn:aws:mediaconnect:ap-northeast-1:111111111111:routerOutput:router-output-id
output "mediaconnect_router_output" {
  value = provider::arn::mediaconnect_router_output("router-output-id")
}
