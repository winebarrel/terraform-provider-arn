# arn:aws:mediaconnect:ap-northeast-1:111111111111:routerInput:router-input-id
output "mediaconnect_router_input" {
  value = provider::arn::mediaconnect_router_input("router-input-id")
}
