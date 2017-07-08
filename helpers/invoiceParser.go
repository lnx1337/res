package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

const numbersRegex = `[0-9]+`

// ParseResponse ayuda a parsear e identificar el tipo de respuesta
// de el servicio facturas
// isLimit cambia a true cuando el resultado de la API responde
// "mas de n resultados"
// number toma el valor de las facturas actuales en el rango de fechas
// number puede ser el limite maximo de resultados por petición
func ParseResponse(resp string) (int64, bool) {

	var isLimit = false
	var re *regexp.Regexp

	re = regexp.MustCompile(numbersRegex)
	numbers := re.FindAllString(resp, -1)[0]

	number, err := strconv.ParseInt(numbers, 10, 64)
	if err != nil {
		panic(err)
	}

	if strings.Contains(resp, "resultados") {
		return number, !isLimit
	}

	return number, isLimit
}
