package transfer

import (
	"github.com/lozovoya/gohomework4_1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc       *card.Service
		ItoICommision int64
		ItoIMin       int64
		ItoECommision int64
		ItoEMin       int64
		EtoICommision int64
		EtoIMin       int64
		EtoECommision int64
		EtoEMin       int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}

	CardSvc := card.NewService("qqq")
	CardSvc.IssueCard("master", 100_000_00, "0000 0000 0000 0000", "rub")
	CardSvc.IssueCard("visa", 15_000_00, "3333 3333 3333 3333", "rub")

	f := fields{&CardSvc, 0,0,0,0,0,0}
	a := args{"000", "000", 10_00}




	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{"ItoI", f fields, a args, 776_00, true },
		//	{"Peguin Bank", []*card.Service{
		//		{0, "visa", 15_000_00, "0000 0000 0000 0000", "rub"},
		//		{1, "master", 15_000_00, "1111 1111 1111 1111", "rub"},
		//	}
		//	}, 0, 0, 0, 0, 0, 0, 0, 0},
		//args({"0000", "0000", 12_000}),
		//},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       tt.fields.CardSvc,
				ItoICommision: tt.fields.ItoICommision,
				ItoIMin:       tt.fields.ItoIMin,
				ItoECommision: tt.fields.ItoECommision,
				ItoEMin:       tt.fields.ItoEMin,
				EtoICommision: tt.fields.EtoICommision,
				EtoIMin:       tt.fields.EtoIMin,
				EtoECommision: tt.fields.EtoECommision,
				EtoEMin:       tt.fields.EtoEMin,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}