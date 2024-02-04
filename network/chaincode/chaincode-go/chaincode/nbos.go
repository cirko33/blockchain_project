package chaincode

import "fmt"

var NationalBankOfSerbia = map[string]ExchangeRate{
	"RSD": {BuyingRate: 1, MiddleRate: 1, SellingRate: 1},
	"USD": {BuyingRate: 107.3524, MiddleRate: 107.6754, SellingRate: 107.9984},
	"AUD": {BuyingRate: 70.7640, MiddleRate: 70.9769, SellingRate: 71.1898},
	"BAM": {BuyingRate: 59.7241, MiddleRate: 59.9038, SellingRate: 60.0835},
	"GBP": {BuyingRate: 136.8441, MiddleRate: 137.2559, SellingRate: 137.6677},
	"DKK": {BuyingRate: 15.6652, MiddleRate: 15.7123, SellingRate: 15.7594},
	"EUR": {BuyingRate: 116.8101, MiddleRate: 117.1616, SellingRate: 117.5131},
	"JPY": {BuyingRate: 0.733133, MiddleRate: 0.735339, SellingRate: 0.737545},
	"CAD": {BuyingRate: 80.2432, MiddleRate: 80.4847, SellingRate: 80.7262},
	"KWD": {BuyingRate: 349.1037, MiddleRate: 350.1542, SellingRate: 351.2047},
	"HUF": {BuyingRate: 0.304804, MiddleRate: 0.305721, SellingRate: 0.306638},
	"NOK": {BuyingRate: 10.2891, MiddleRate: 10.3201, SellingRate: 10.3511},
	"PLN": {BuyingRate: 27.0506, MiddleRate: 27.1320, SellingRate: 27.2134},
	"RUB": {BuyingRate: 1.1845, MiddleRate: 1.1881, SellingRate: 1.1917},
	"CZK": {BuyingRate: 4.6940, MiddleRate: 4.7081, SellingRate: 4.7222},
	"CHF": {BuyingRate: 125.1850, MiddleRate: 125.5617, SellingRate: 125.9384},
	"SEK": {BuyingRate: 10.3339, MiddleRate: 10.3650, SellingRate: 10.3961},
}

func ConvertCurrency(amount float64, fromIso, toIso string) (float64, error) {

	fromRate, err := GetExchangeRate(fromIso)
	if err != nil {
		return 0, err
	}

	toRate, err1 := GetExchangeRate(toIso)
	if err1 != nil {
		return 0, err1
	}

	exchangeRate := fromRate.MiddleRate / toRate.MiddleRate

	return amount * exchangeRate, nil
}

func GetExchangeRate(iso string) (*ExchangeRate, error) {
	rate, exists := NationalBankOfSerbia[iso]
	if !exists {
		return nil, fmt.Errorf("The exchange rate for %s not found", iso)
	}
	return &rate, nil
}