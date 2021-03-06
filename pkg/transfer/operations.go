package transfer

import (
	"github.com/lozovoya/gohomework4_1/pkg/card"
)

type Service struct {
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

func NewService(cardSvc *card.Service, itoICommision int64, itoIMin int64, itoECommision int64, itoEMin int64,
	etoICommision int64, etoIMin int64, etoECommision int64, etoEMin int64) *Service {
	return &Service{CardSvc: cardSvc,
		ItoICommision: itoICommision, ItoIMin: itoIMin,
		ItoECommision: itoECommision, ItoEMin: itoEMin,
		EtoICommision: etoICommision, EtoIMin: etoIMin,
		EtoECommision: etoECommision, EtoEMin: etoEMin}
}

func (s *Service) Card2Card(from string, to string, amount int64) (total int64, ok bool) {
	fromCard := s.CardSvc.SearchByNumber(from)
	toCard := s.CardSvc.SearchByNumber(to)

	// I to I
	if (fromCard != nil) && (toCard != nil) {
		if fromCard.Balance < amount {
			return amount, false
		}
		fromCard.Balance -= amount
		toCard.Balance += amount
		return amount, true

	}

	// I to E
	if (fromCard != nil) && (toCard == nil) {
		commission := amount * s.ItoECommision / 1000
		total := amount + commission
		if fromCard.Balance < total {
			return total, false
		}
		fromCard.Balance -= total
		return total, true

	}

	// E to I
	if (fromCard == nil) && (toCard != nil) {
		toCard.Balance += amount
		return amount, true
	}

	// E to E
	if (fromCard == nil) && (toCard == nil) {
		commission := amount * s.EtoECommision / 1000
		if commission > s.EtoEMin {
			total = amount + commission
			return total, true
		}
		total = amount + s.EtoEMin
		return total, true

	}
	return 0, false
}
