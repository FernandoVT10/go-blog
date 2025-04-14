package utils

import (
    "fmt"
    "time"
)

func GetTimeAgo(thenTime time.Time) string {
    nowTime := time.Now()

    var units = [...]string{"year", "month", "day", "hour", "minute"}

    for _, unit := range units {
        var now, then int

        switch(unit) {
        case "year":
            now = nowTime.Year()
            then = thenTime.Year()
        case "month":
            now = int(nowTime.Month())
            then = int(thenTime.Month())
        case "day":
            now = nowTime.Day()
            then = thenTime.Day()
        case "hour":
            now = nowTime.Hour()
            then = thenTime.Hour()
        case "minute":
            now = nowTime.Minute()
            then = thenTime.Minute()
        }

        if now <= then { continue }
        diff := now - then

        if diff > 1 {
            // here we make unit plural
            return fmt.Sprintf("%d %ss ago", diff, unit)
        } else {
            return fmt.Sprintf("%d %s ago", diff, unit)
        }

    }

    return "A moment ago"
}

func FormatTime(t time.Time) string {
    day := t.Day()
    month := t.Month().String()
    year := t.Year()

    return fmt.Sprintf("%d %s %d", day, month, year)
}
