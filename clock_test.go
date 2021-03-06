package clockgsl
import (
  "strings"
  "strconv"
  "testing"
  "fmt"
)

func timeToDayFraction(timestamp string) (float64, error) {
	timesplit := strings.Split(timestamp, ":")
	if len(timesplit) != 3 {
		return 0.0, fmt.Errorf("Time needs to be in format HH:MM:SS but was %s", timestamp)
	}
	hours, err := strconv.Atoi(timesplit[0])
	if err != nil {
		return 0.0, fmt.Errorf("Time needs to be in format HH:MM:SS but was %s", timestamp)
	}
	mins, err := strconv.Atoi(timesplit[1])
	if err != nil {
		return 0.0, fmt.Errorf("Time needs to be in format HH:MM:SS but was %s", timestamp)
	}
	secs, err := strconv.Atoi(timesplit[2])
	if err != nil {
		return 0.0, fmt.Errorf("Time needs to be in format HH:MM:SS but was %s", timestamp)
	}
	secondsSinceMidnight := float64(hours * 60 * 60 + mins * 60 + secs)
	return secondsSinceMidnight / float64(1.0 * 24 * 60 * 60), nil
}

func TestLocaltime(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "00:00:00" },
		{ "23:59:59", "23:59:59" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getLocaltime(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestLocaldate(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getLocaldate(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestHolocenedate(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getHolocenedate(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestUtc(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "00:00:00 UT" },
		{ "23:59:59", "23:59:59 UT" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getUtc(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestBeattime(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "@000" },
		{ "23:59:59", "@999" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getBeattime(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestHextime(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "0_00_0" },
		{ "02:59:59", "1_FF_F" },
		{ "03:00:00", "2_00_0" },
		{ "09:30:00", "6_55_5" },
		{ "14:00:00", "9_55_5" },
		{ "17:30:00", "B_AA_A" },
		{ "23:59:59", "F_FF_F" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getHextime(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestMetric(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "00.000 LMT" },
		{ "00:34:56", "02.425 LMT" },
		{ "13:51:40", "57.754 LMT" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getMetric(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestBinaryhours(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "000000" },
		{ "01:00:00", "000001" },
		{ "12:00:00", "001100" },
		{ "23:59:59", "010111" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getBinaryhours(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestBinarymins(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "000000" },
		{ "00:01:00", "000001" },
		{ "11:59:59", "111011" },
		{ "12:00:00", "000000" },
		{ "23:59:59", "111011" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getBinarymins(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestBinarysecs(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "000000" },
		{ "00:00:01", "000001" },
		{ "00:00:31", "011111" },
		{ "00:00:32", "100000" },
		{ "00:00:47", "101111" },
		{ "00:00:48", "110000" },
		{ "00:00:02", "000010" },
		{ "00:00:59", "111011" },
		{ "00:01:00", "000000" },
		{ "00:59:59", "111011" },
		{ "05:59:59", "111011" },
		{ "11:59:59", "111011" },
		{ "12:00:00", "000000" },
		{ "23:59:59", "111011" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getBinarysecs(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestBells(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "Eight bells" },
		{ "00:29:59", "Eight bells" },
		{ "00:30:00", "One bell" },
		{ "18:29:59", "Four bells" },
		{ "18:30:00", "One bell" },
		{ "19:59:59", "Three bells" },
		{ "20:00:00", "Eight bells" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getBells(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
func TestWatch(t *testing.T) {
	testdata := []struct {
		in string
		out string
	}{
		{ "00:00:00", "Middle watch" },
		{ "16:45:00", "First dog watch" },
		{ "23:59:59", "First watch" },
	}
        for _, tt := range testdata {
                fraction, err := timeToDayFraction(tt.in)
				if err != nil {
                        t.Errorf("Error: %s", err)
				}
                s := getWatch(fraction)
                if s != tt.out {
                        t.Errorf("%q => %q, want %q (fraction: %q)", tt.in, s, tt.out, fraction)
                }
        }

}
