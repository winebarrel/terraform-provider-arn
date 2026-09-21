// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: chatbot
// Source: https://servicereference.us-east-1.amazonaws.com/v1/chatbot/chatbot.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "chatbot_chatbot_configuration", Service: "chatbot", Resource: "ChatbotConfiguration", Template: "arn:${Partition}:chatbot::${Account}:chat-configuration/${ConfigurationType}/${ChatbotConfigurationName}"},
		{Name: "chatbot_custom_action", Service: "chatbot", Resource: "custom-action", Template: "arn:${Partition}:chatbot::${Account}:custom-action/${ActionName}"},
	})
}
