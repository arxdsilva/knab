package domains

import "testing"

func TestAccount_Verify(t *testing.T) {
	tests := []struct {
		name           string
		DocumentNumber string
		wantErr        bool
	}{
		{
			name:           "ok",
			DocumentNumber: "1234",
			wantErr:        false,
		},
		{
			name:           "err not a number",
			DocumentNumber: "abcde",
			wantErr:        true,
		},
		{
			name:           "err not a number",
			DocumentNumber: "abc1234de",
			wantErr:        true,
		},
		{
			name:           "err empty doc number",
			DocumentNumber: "",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{DocumentNumber: tt.DocumentNumber}
			if err := a.Verify(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
