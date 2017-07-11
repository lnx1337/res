package utils

import (
	"fmt"
	"time"

	"github.com/metakeule/fmtdate"
)

/*
* @function GetMidDate Obtiene la fecha intermedia entre dos fechas
* @var startDate fecha de inicio en formato YYYY-MM-DD
* @var endDate fecha final en formato YYYY-MM-DD
 */
func GetMidDate(startDate, endDate string) string {

	dSParsed := Parse(startDate)
	dEParsed := Parse(endDate)

	daysDS := GetDateInDays(dSParsed)
	daysDE := GetDateInDays(dEParsed)

	toInt := (daysDE - daysDS) / 2
	dateToParse := dSParsed.AddDate(0, 0, toInt)

	midDate := GetDate(dateToParse)

	return midDate
}

/*
* @function GetDate retorna fecha en formato YYYY-MM-DD
* @var fecha en formato time.Time
 */
func GetDate(date time.Time) string {
	year := date.Year()

	monthNumber := int(date.Month())
	month := GetMonth(monthNumber)

	dayNumber := date.Day()
	days := GetDays(dayNumber)

	return fmt.Sprint(year, "-", month, "-", days)
}

/*
*@function GetMonth retorna mes valido en formato MM
*@var month numero de mes int
 */
func GetMonth(month int) string {
	if month < 10 {
		return fmt.Sprint("0", month)
	}
	return fmt.Sprint(month)
}

/*
*@function GetDays retorna día valido en formato DD
*@var numero de día en formato int
 */
func GetDays(days int) string {
	if days < 10 {
		return fmt.Sprint("0", days)
	}
	return fmt.Sprint(days)
}

/*
* @function Parse permite manejar las fechas en go
* @var date recibe una fecha en formato YYYY-MM-DD
* @return fecha para poder hacer operaciones con go
 */
func Parse(date string) time.Time {
	parsed, _ := fmtdate.Parse("YYYY-MM-DD", date)
	return parsed
}

/*
* @fuction GetDateInDays
* @ret fecha en dias en formato int
 */
func GetDateInDays(date time.Time) int {
	days := HoursToDays(DateToHours(date))
	return days
}

/*
* @fuction DateToHours
* @return Fecha recibida en horas desde 1900-01-01
 */
func DateToHours(date time.Time) int {
	dateBASE, _ := fmtdate.Parse("YYYY-MM-DD", "1900-01-01")
	dDS := date.Sub(dateBASE)
	hours := int(dDS.Hours())
	return hours
}

/* @function HoursToDays
* @return horas en días
 */
func HoursToDays(hours int) int {
	return (hours / 24)
}
