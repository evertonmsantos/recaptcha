package recaptcha

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/evertonmsantos/request"
)

func Invisible(endpoint string) (string, error) {

	req, err := request.Get(endpoint, make(map[string]string), false)

	if err != nil {
		return "", err
	}

	inputToken := regexp.MustCompile(`(?m)<input type="hidden" id="recaptcha-token" value="(.*?)"`).FindStringSubmatch(req.Body)

	if len(inputToken) == 0 {
		return "", errors.New("input token não localizado")
	}

	params, err := url.ParseQuery(endpoint)

	if err != nil {
		return "", errors.New("erro ao decodificar os parametros")
	}

	req, err = request.Post(fmt.Sprintf(`https://www.google.com/recaptcha/api2/reload?k=%s`, params["k"][0]), fmt.Sprintf(`c=%s&reason=q&k=%s`, inputToken[1], params["k"][0]), map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, false)

	if err != nil {
		return "", err
	}

	token := regexp.MustCompile(`(?m)"rresp","(.*?)"`).FindStringSubmatch(req.Body)

	if len(token) == 0 {
		return "", errors.New("não foi possivel decifrar o captcha")
	}

	return token[1], nil

}
