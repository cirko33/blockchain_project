package models

type Bank struct {
	Id           string             `json:"id"`
	HeadQuarters string             `json:"headQuarters"`
	YearFounded  int                `json:"yearFounded"`
	PIB          string             `json:"pib"`
	Users        map[string]Person  `json:"users"`
	Accounts     map[string]Account `json:"accounts"`
}
