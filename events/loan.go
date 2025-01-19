package events

import (
	"github.com/emacsified/r2d2-trader/models"
	"time"
)

const T_PURCHASED = "PURCHASED"
const T_PAID = "PAID"

type Loan struct {
	name string
	time time.Time

	Type    string
	Account models.Account
}

func (e Loan) New(args ...interface{}) Event {
	return &Loan{
		name:    "LOAN",
		time:    time.Now(),
		Type:    args[0].(string),
		Account: args[1].(models.Account),
	}
}

func (e *Loan) Name() string {
	return e.name
}

func (e *Loan) Time() time.Time {
	return e.time
}
