# arn:aws:appsync:ap-northeast-1:111111111111:apis/api-id/channelNamespace/channel-namespace-name
output "appsync_channel_namespace" {
  value = provider::arn::appsync_channel_namespace("api-id", "channel-namespace-name")
}
