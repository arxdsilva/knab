package domains

import "testing"

func TestAccount_Verify(t *testing.T) {
	tests := []struct {
		name           string
		DocumentNumber string
		TotalCredit    float64
		wantErr        bool
	}{
		{
			name:           "ok",
			DocumentNumber: "1234",
			wantErr:        false,
			TotalCredit:    1.0,
		},
		{
			name:           "err not a number",
			DocumentNumber: "abcde",
			wantErr:        true,
			TotalCredit:    1.0,
		},
		{
			name:           "err not a number",
			DocumentNumber: "abc1234de",
			wantErr:        true,
			TotalCredit:    1.0,
		},
		{
			name:           "err empty doc number",
			DocumentNumber: "",
			wantErr:        true,
			TotalCredit:    1.0,
		},
		{
			name:           "negative total credit",
			DocumentNumber: "12313131",
			wantErr:        true,
			TotalCredit:    -10.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{DocumentNumber: tt.DocumentNumber, TotalCredit: tt.TotalCredit}
			if err := a.Verify(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
