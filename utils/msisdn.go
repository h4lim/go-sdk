package utils

import "errors"

var InvalidMsIsdnError = errors.New("invalid msisdn number")

func MsisdnToLocalID(msisdn string) (string, error) {
	if len(msisdn) < 4 || msisdn[0:3] != "+62" {
		return "", InvalidMsIsdnError
	}

	local := msisdn[3:]
	if local[0] == '0' {
		return "", InvalidMsIsdnError
	}
	local = "0" + local

	if local[1] != '8' && local[1:4] != "999" {
		return "", InvalidMsIsdnError
	}

	for _, c := range local {
		if c < '0' || c > '9' {
			return "", InvalidMsIsdnError
		}
	}
	return local, nil
}

func MsisdnToLocalIDReformat(msisdn string) (string, error) {
	if len(msisdn) < 4 || msisdn[0:3] != "+62" {
		return "", InvalidMsIsdnError
	}

	local := msisdn[3:]
	if local[0] == '0' {
		return "", InvalidMsIsdnError
	}

	if len(local) == 9 {
		local = "000" + local
	} else if len(local) == 10 {
		local = "00" + local
	} else if len(local) == 11 {
		local = "0" + local
	} else if len(local) == 12 {
		local = local
	} else {
		local = "0" + local
	}

	return local, nil
}
