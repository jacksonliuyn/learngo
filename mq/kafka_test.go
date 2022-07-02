package mq

import "testing"

func TestProducer(t *testing.T) {
	type args struct {
		topic string
		limit int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test", args{"test", 10}},
		{"test", args{"test", 10}},
		{"test", args{"test", 10}},
		{"test", args{"test", 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Producer(tt.args.topic, tt.args.limit)
		})
	}
}
