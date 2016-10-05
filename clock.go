package clockgsl
import (
  "time"
  "fmt"
  "math"
)

type supplier interface {
	FractionOfDayPassed() float64
}

func getLocaltime(fractionOfDayPassed float64) string {
	hourten := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 5/12)
	hourunit := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/24)
	minuteten := getSixdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/144)
	minuteunit := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/1440)
	secondten := getSixdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/8640)
	secondunit := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/86400)
    return fmt.Sprintf("%s%s:%s%s:%s%s" , hourten, hourunit, minuteten, minuteunit, secondten, secondunit)
}
func getLocaldate(fractionOfDayPassed float64) string {
    Year := time.Now().Year()
    Month := time.Now().Month()
    Day := time.Now().Day()
    return fmt.Sprintf("%d-%02d-%02d" , Year, Month, Day)
}
func getUtc(fractionOfDayPassed float64) string {
	hourten := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 5/12)
	hourunit := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/24)
	minuteten := getSixdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/144)
	minuteunit := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/1440)
	secondten := getSixdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/8640)
	secondunit := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/86400)
    return fmt.Sprintf("%s%s:%s%s:%s%s UT" , hourten, hourunit, minuteten, minuteunit, secondten, secondunit)
}
func getBeattime(fractionOfDayPassed float64) string {
	hundreds := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/10)
	tens := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/100)
	units := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/1000)
    return fmt.Sprintf("@%s%s%s" , hundreds, tens, units)
}
func getHextime(fractionOfDayPassed float64) string {
	hexhours := getHexdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/16)
	hexminutetens := getHexdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/256)
	hexminuteunits := getHexdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/4096)
	hexseconds := getHexdigitsOutputElement(fractionOfDayPassed, 1.0 * 1/65536)
    return fmt.Sprintf("%s_%s%s_%s" , hexhours, hexminutetens, hexminuteunits, hexseconds)
}
func getMetric(fractionOfDayPassed float64) string {
	decidays := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/10)
	centidays := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/100)
	millidays := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/1000)
	quadrodays := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/10000)
	quintodays := getDigitsOutputElement(fractionOfDayPassed, 1.0 * 1/100000)
    return fmt.Sprintf("%s%s.%s%s%s LMT" , decidays, centidays, millidays, quadrodays, quintodays)
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
	out = out + fmt.Sprintln(getBells(localSupplier().FractionOfDayPassed()))
	out = out + fmt.Sprintln(getWatch(localSupplier().FractionOfDayPassed()))
    return out
}

func GetAllAsJSON() string {
	return fmt.Sprintf(`{"localTime":"%s","localDate":"%s","utc":"%s","beatTime":"%s","hexTime":"%s","metric":"%s","bells":"%s","watch":"%s","tail": "true"}`, getLocaltime(localSupplier().FractionOfDayPassed()), getLocaldate(localSupplier().FractionOfDayPassed()), getUtc(utcSupplier().FractionOfDayPassed()), getBeattime(bmtSupplier().FractionOfDayPassed()), getHextime(localSupplier().FractionOfDayPassed()), getMetric(localSupplier().FractionOfDayPassed()), getBells(localSupplier().FractionOfDayPassed()), getWatch(localSupplier().FractionOfDayPassed()))
}

