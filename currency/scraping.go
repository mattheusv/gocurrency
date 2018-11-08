package currency

// Currency struct
type Currency struct {
	Name             string `json:"currency_name"`
	Value            string `json:"currency_value"`
	DateVerification string `json:"verification_date"`
}

const (
	//TRADITIONAL ...
	TRADITIONAL = 1
	//DIGITAL ...
	DIGITAL = 0
)

//GetCurrrency ...
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
		return Currency{}, nil
	}
}

//GetAllCurrency ...
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
		return nil, nil
	}
}
