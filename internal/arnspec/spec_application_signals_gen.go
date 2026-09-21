// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: application-signals
// Source: https://servicereference.us-east-1.amazonaws.com/v1/application-signals/application-signals.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "application_signals_instrumentation_config", Service: "application-signals", Resource: "instrumentationConfig", Template: "arn:${Partition}:application-signals:${Region}:${Account}:instrumentationConfig/${Service}/${Environment}/${SignalType}/${LocationHash}"},
		{Name: "application_signals_slo", Service: "application-signals", Resource: "slo", Template: "arn:${Partition}:application-signals:${Region}:${Account}:slo/${SloName}"},
	})
}
