package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

const findNumbersRegex = `[0-9]+`

/* ParseResponse
* @function ParseResponse
* @returns int, bool
* @public
* @var iLimit cambia a true cuando el resultado de la API responde
* "mas de n resultados"
* @var number toma el valor de las facturas actuales en el rango de fechas
* number puede ser el limite maximo de resultados por petición
 */

// ParseResponse Parsea e identifica el tipo de respuesta de el servicio facturas
func ParseResponse(resp string) (int64, bool) {

	var isLimit = false
	var re *regexp.Regexp

	re = regexp.MustCompile(findNumbersRegex)
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
