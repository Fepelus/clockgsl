package clockgsl
import (
  "time"
  "fmt"
  "math"
)

type supplier interface {
	FractionOfDayPassed() float64
}

func floatingModulo(fractionOfDayPassed float64, resetFraction float64) float64 {
	division := math.Floor(fractionOfDayPassed / resetFraction )
	return (fractionOfDayPassed - (division * resetFraction))
}


func getLocaltime(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1.0
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	hourten := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 5/12)
	hourunit := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/24)
	minuteten := getSixdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/144)
	minuteunit := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/1440)
	secondten := getSixdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/8640)
	secondunit := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/86400)
    return fmt.Sprintf("%s%s:%s%s:%s%s" , hourten, hourunit, minuteten, minuteunit, secondten, secondunit)
}
func getLocaldate(fractionOfDayPassed float64) string {
    Year := time.Now().Year()
    Month := time.Now().Month()
    Day := time.Now().Day()
    return fmt.Sprintf("%d-%02d-%02d" , Year, Month, Day)
}
func getUtc(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1.0
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	hourten := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 5/12)
	hourunit := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/24)
	minuteten := getSixdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/144)
	minuteunit := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/1440)
	secondten := getSixdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/8640)
	secondunit := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/86400)
    return fmt.Sprintf("%s%s:%s%s:%s%s UT" , hourten, hourunit, minuteten, minuteunit, secondten, secondunit)
}
func getBeattime(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1.0
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	hundreds := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/10)
	tens := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/100)
	units := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/1000)
    return fmt.Sprintf("@%s%s%s" , hundreds, tens, units)
}
func getHextime(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1.0
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	hexhours := getHexdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/16)
	hexminutetens := getHexdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/256)
	hexminuteunits := getHexdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/4096)
	hexseconds := getHexdigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/65536)
    return fmt.Sprintf("%s_%s%s_%s" , hexhours, hexminutetens, hexminuteunits, hexseconds)
}
func getMetric(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1.0
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	decidays := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/10)
	centidays := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/100)
	millidays := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/1000)
	hectomicrodays := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/10000)
	dekamicrodays := getDigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/100000)
    return fmt.Sprintf("%s%s.%s%s%s LMT" , decidays, centidays, millidays, hectomicrodays, dekamicrodays)
}
func getBinaryhours(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1.0
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	bhsixt := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 2/3)
	bheight := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/3)
	bhfour := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/6)
	bhtwo := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/12)
	bhunit := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/24)
    return fmt.Sprintf("0%s%s%s%s%s" , bhsixt, bheight, bhfour, bhtwo, bhunit)
}
func getBinarymins(fractionOfDayPassed float64) string {
    resetFraction := 1.0 * 1/24
   
    resetFractionOfDayPassed := floatingModulo(fractionOfDayPassed, resetFraction)
	bhthirt := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/45)
	bhsixt := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/90)
	bheight := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/180)
	bhfour := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/360)
	bhtwo := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/720)
	bhunit := getBinarydigitsOutputElement(resetFractionOfDayPassed, 1.0 * 1/1440)
    return fmt.Sprintf("%s%s%s%s%s%s" , bhthirt, bhsixt, bheight, bhfour, bhtwo, bhunit)
}
func getBells(fractionOfDayPassed float64) string {
    bellcollection := []string{"Eight bells","One bell","Two bells","Three bells","Four bells","Five bells","Six bells","Seven bells","Eight bells","One bell","Two bells","Three bells","Four bells","Five bells","Six bells","Seven bells","Eight bells","One bell","Two bells","Three bells","Four bells","Five bells","Six bells","Seven bells","Eight bells","One bell","Two bells","Three bells","Four bells","Five bells","Six bells","Seven bells","Eight bells","One bell","Two bells","Three bells","Four bells","One bell","Two bells","Three bells","Eight bells","One bell","Two bells","Three bells","Four bells","Five bells","Six bells","Seven bells",  }
    bellidx := int(math.Mod(math.Floor((fractionOfDayPassed) / (1.0 * 1/48)),float64(len(bellcollection))))
    bell := bellcollection[bellidx]
    return fmt.Sprintf("%s" , bell)
}
func getWatch(fractionOfDayPassed float64) string {
    watchcollection := []string{"Middle","Middle","Morning","Morning","Forenoon","Forenoon","Afternoon","Afternoon","First dog","Last dog","First","First",  }
    watchidx := int(math.Mod(math.Floor((fractionOfDayPassed) / (1.0 * 1/12)),float64(len(watchcollection))))
    watch := watchcollection[watchidx]
    return fmt.Sprintf("%s watch" , watch)
}


func getDigitsOutputElement(fractionOfDayPassed float64, dayFraction float64) string {
    collection := []string{"0","1","2","3","4","5","6","7","8","9",}
    idx := int(math.Mod(math.Floor((fractionOfDayPassed) / dayFraction),float64(len(collection))))
    return collection[idx]
}
func getSixdigitsOutputElement(fractionOfDayPassed float64, dayFraction float64) string {
    collection := []string{"0","1","2","3","4","5",}
    idx := int(math.Mod(math.Floor((fractionOfDayPassed) / dayFraction),float64(len(collection))))
    return collection[idx]
}
func getHexdigitsOutputElement(fractionOfDayPassed float64, dayFraction float64) string {
    collection := []string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F",}
    idx := int(math.Mod(math.Floor((fractionOfDayPassed) / dayFraction),float64(len(collection))))
    return collection[idx]
}
func getBinarydigitsOutputElement(fractionOfDayPassed float64, dayFraction float64) string {
    collection := []string{"0","1",}
    idx := int(math.Mod(math.Floor((fractionOfDayPassed) / dayFraction),float64(len(collection))))
    return collection[idx]
}

func localSupplier() supplier {
	return thesupplier { time.Now() }
}
func utcSupplier() supplier {
	location := time.FixedZone("utc", 0 * 60 * 60)
	return thesupplier{ time.Now().In(location) }
}
func bmtSupplier() supplier {
	location := time.FixedZone("bmt", +1 * 60 * 60)
	return thesupplier{ time.Now().In(location) }
}


type thesupplier struct {
	clock time.Time
}
func (self thesupplier) FractionOfDayPassed() float64 {
	nanoSecondsPerDay := float64(24 * 60 * 60 * 1000 * 1000 * 1000)
	round := self.clock.Round(time.Nanosecond)
	midnight := round.Add(-time.Duration(round.Hour()) * time.Hour).Add(-time.Duration(round.Minute()) * time.Minute).Add(-time.Duration(round.Second()) * time.Second).Add(-time.Duration(round.Nanosecond()) * time.Nanosecond)
	return float64(self.clock.Sub(midnight).Nanoseconds()) / nanoSecondsPerDay
}

func GetAllAsString() string {
    out := ""
	out = out + fmt.Sprintln(getLocaltime(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getLocaldate(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getUtc(utcSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getBeattime(bmtSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getHextime(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getMetric(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getBinaryhours(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getBinarymins(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getBells(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getWatch(localSupplier().FractionOfDayPassed()))
    return out
}

func GetAllAsJSON() string {
    return fmt.Sprintf(
        `{"localTime":"%s","localDate":"%s","utc":"%s","beatTime":"%s","hexTime":"%s","metric":"%s","binaryhours":"%s","binarymins":"%s","bells":"%s","watch":"%s","tail": "true"}`,
		getLocaltime(localSupplier().FractionOfDayPassed()),
		getLocaldate(localSupplier().FractionOfDayPassed()),
		getUtc(utcSupplier().FractionOfDayPassed()),
		getBeattime(bmtSupplier().FractionOfDayPassed()),
		getHextime(localSupplier().FractionOfDayPassed()),
		getMetric(localSupplier().FractionOfDayPassed()),
		getBinaryhours(localSupplier().FractionOfDayPassed()),
		getBinarymins(localSupplier().FractionOfDayPassed()),
		getBells(localSupplier().FractionOfDayPassed()),
		getWatch(localSupplier().FractionOfDayPassed()),
   )
}

