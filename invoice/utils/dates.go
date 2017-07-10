package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/metakeule/fmtdate"
)

/*
* @function GetMidDate Obtiene la fecha intermedia entre dos fechas
* @var startDate fecha de inicio en formato YYYY-MM-DD
* @var endDate fecha final en formato YYYY-MM-DD
 */
func GetMidDate(startDate, endDate string) string {
	var day string
	var year string
	var month string

	dSParsed := Parse(startDate)
	dEParsed := Parse(endDate)

	daysDS := GetDateInDays(dSParsed)
	daysDE := GetDateInDays(dEParsed)

	// total de dias
	dateMidInDays := float64(((daysDE - daysDS) / 2) + daysDS)

	// Years
	yy := (dateMidInDays / 365.02)
	onlyY := int(yy)
	yearDec := yy - float64(onlyY)

	if dSParsed.Year() == dEParsed.Year() {
		year = fmt.Sprint(dSParsed.Year())
	} else if onlyY >= 100 {
		yyyy := onlyY - 100
		year = fmt.Sprint("20", yyyy)
		if yyyy < 10 {
			year = fmt.Sprint("200", yyyy)
		}
	} else {
		year = fmt.Sprint("19", onlyY)
		if onlyY < 10 {
			year = fmt.Sprint("190", onlyY)
		}
	}

	// Month
	mm := yearDec * 365.20 / 30.42
	onlyM := int(mm)
	monthDec := mm - float64(onlyM)

	if onlyM < 10 {
		month = fmt.Sprint("0", onlyM)
		if onlyM == 0 {
			month = "12"
		}
	} else {
		month = strconv.Itoa(onlyM)
	}

	// Days
	dd := int(monthDec * 30.42)
	if dd < 10 {
		day = fmt.Sprint("0", dd)
	} else {
		day = strconv.Itoa(dd)
	}

	if dd < 1 {
		day = "01"
	}

	return fmt.Sprint(year, "-", month, "-", day)
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
