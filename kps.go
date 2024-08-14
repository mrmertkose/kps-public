package kps

import (
	"errors"
	"github.com/hooklift/gowsdl/soap"
	"github.com/mrmertkose/kps-public/service"
)

const baseUrl = "https://tckimlik.nvi.gov.tr/Service/KPSPublic.asmx?WSDL"

func GetService(identityNumber int64, name string, surname string, year int32) (bool, error) {

	if identityNumber == 0 || name == "" || surname == "" || year == 0 {
		return false, errors.New("information cannot be left empty")
	}

	client := soap.NewClient(baseUrl)

	conn := service.NewKPSPublicSoap(client)

	cfg := &service.TCKimlikNoDogrula{
		TCKimlikNo: identityNumber,
		Ad:         name,
		Soyad:      surname,
		DogumYili:  year,
	}

	check, err := conn.TCKimlikNoDogrula(cfg)
	if err != nil {
		return false, err
	}

	return check.TCKimlikNoDogrulaResult, nil
}
