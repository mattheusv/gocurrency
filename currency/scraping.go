package currency

import (
	"fmt"

	"github.com/msalcantara/gocurrency/errors"
)

// Currency struct
type Currency struct {
	Name             string `json:"currency_name"`
	Value            string `json:"currency_value"`
	DateVerification string `json:"verification_date"`
}

const (
	//TRADITIONAL traditional currencys
	TRADITIONAL = 1
	//DIGITAL digital currencys
	DIGITAL = 0
)

//GetCurrrency get specific currecy by name and type
//type = TRADITIONAL or DIGITAL
func GetCurrrency(name string, t int) (Currency, error) {
	switch t {
	case TRADITIONAL:
		trad := traditional{}
		c, err := trad.GetCurrrency(name)
		if err != nil {
			return Currency{}, err
		}
		return c, nil
	case DIGITAL:
		d := digital{}
		c, err := d.GetCurrrency(name)
		if err != nil {
			return Currency{}, err
		}
		return c, nil
	default:
		return Currency{}, errors.New(fmt.Sprintf("invalid type of currency %d", t))
	}
}

//GetAllCurrency get all currecys by type
//type = TRADITIONAL or DIGITAL
func GetAllCurrency(t int) ([]Currency, error) {
	switch t {
	case TRADITIONAL:
		trad := traditional{}
		c, err := trad.GetAll()
		if err != nil {
			return nil, err
		}
		return c, nil
	case DIGITAL:
		d := digital{}
		c, err := d.GetAll()
		if err != nil {
			return nil, err
		}
		return c, nil
	default:
		return nil, errors.New(fmt.Sprintf("invalid type of currency %d", t))
	}
}
