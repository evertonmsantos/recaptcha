package recaptcha

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/evertonmsantos/request"
)

func Verify(endpoint, token string) (string, error) {

	params, err := url.ParseQuery(endpoint)

	if err != nil {
		return "", errors.New("erro ao decodificar os parametros")
	}

	req, err := request.Post(fmt.Sprintf(`https://www.google.com/recaptcha/api2/userverify?k=%s`, params["k"][0]), fmt.Sprintf(`v=%s&c=%s&response=eyJyZXNwb25zZSI6IiIsInMiOiIzZmEzIiwiZSI6ImJXMXZDS2cifQ..&t=4867&ct=4867`, params["v"][0], token), map[string]string{"content-type": "application/x-www-form-urlencoded;charset=UTF-8"}, false)

	if err != nil {
		return "", err
	}

	utoken := regexp.MustCompile(`(?m)"uvresp","(.*?)"`).FindStringSubmatch(req.Body)

	if len(utoken) == 0 {
		return "", errors.New("não foi possivel decifrar o captcha")
	}

	return utoken[1], nil

}
