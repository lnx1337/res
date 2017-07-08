package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/metakeule/fmtdate"
)

//GetMidDate Obtiene la fecha intermedia entre dos fechas
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
	yearMod := yy - float64(onlyY)

	if dSParsed.Year() == dEParsed.Year() {
		year = fmt.Sprint(dSParsed.Year())
	}

	// Month
	mm := yearMod * 365.20 / 30.42
	onlyM := int(mm)
	monthMod := mm - float64(onlyM)

	if onlyM < 10 {
		month = fmt.Sprint("0", onlyM)
	} else {
		month = strconv.Itoa(onlyM)
	}

	// Days
	dd := int(monthMod * 30.42)

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

// Parse permite manejar las fechas en go
func Parse(date string) time.Time {
	parsed, _ := fmtdate.Parse("YYYY-MM-DD", date)
	return parsed
}

// GetDateInDays retorna la fecha en dias
func GetDateInDays(date time.Time) int {
	days := HoursToDays(DateToHours(date))
	return days
}

// DateToHours retorna la fecha enviada en horas desde 1900-01-01
func DateToHours(date time.Time) int {
	dateBASE, _ := fmtdate.Parse("YYYY-MM-DD", "1900-01-01")
	dDS := date.Sub(dateBASE)
	hours := int(dDS.Hours())
	return hours
}

// HoursToDays Esta función convierte las horas en días
func HoursToDays(hours int) int {
	return (hours / 24)
}
