package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_Verify(t *testing.T) {
	tests := []struct {
		name    string
		tr      Transaction
		wantErr bool
	}{
		{
			name: "ok",
			tr: Transaction{
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          1.0,
			},
			wantErr: false,
		},
		{
			name: "invalid account ID",
			tr: Transaction{
				OperationTypeID: 1,
				Amount:          1.0,
			},
			wantErr: true,
		},
		{
			name: "invalid operation ID",
			tr: Transaction{
				AccountID:       1,
				OperationTypeID: 10,
				Amount:          1.0,
			},
			wantErr: true,
		},
		{
			name: "invalid amount",
			tr: Transaction{
				AccountID:       1,
				OperationTypeID: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Verify(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_Verify_Amount_Negative(t *testing.T) {
	tr := Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          10,
	}
	err := tr.Verify()
	assert.Nil(t, err)
	assert.Equal(t, float64(-10), tr.Amount)
}

func TestTransaction_Verify_Amount_Positive(t *testing.T) {
	tr := Transaction{
		AccountID:       1,
		OperationTypeID: 4,
		Amount:          10,
	}
	err := tr.Verify()
	assert.Nil(t, err)
	assert.Equal(t, float64(10), tr.Amount)
}
