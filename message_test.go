package gotestlib

import "testing"

func TestMessage_Hello(t *testing.T) {
	type fields struct {
		Title   string
		Content string
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{name: "my test",
			fields: fields{Title: "my title", Content: "my content"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Message{
				Title:   tt.fields.Title,
				Content: tt.fields.Content,
			}
			this.Hello()
		})
	}
}
