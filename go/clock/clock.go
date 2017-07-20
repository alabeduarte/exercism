package clock

import (
  "fmt"
  "strconv"
  "time"
)

const testVersion = 4

type Clock struct {
  Hour int
  Minute int
}

func NewTime() time.Time {
  t := time.Now()

  return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func ConvertToHoursDuration(hour int) time.Duration {
  return time.Duration(hour) * time.Hour
}

func ConvertToMinutesDuration(minute int) time.Duration {
  return time.Duration(minute) * time.Minute
}

func New(hour, minute int) Clock {
  durationHours := ConvertToHoursDuration(hour)
  durationMinutes := ConvertToMinutesDuration(minute)
  duration := time.Duration(durationHours.Minutes() + durationMinutes.Minutes())

  clock := NewTime().Add(duration * time.Minute)

  return Clock{clock.Hour(), clock.Minute()}
}

func (clock Clock) Add(minute int) Clock {
  return New(clock.Hour, clock.Minute + minute)
}

func (clock Clock) String() string {
  hour, minute := strconv.Itoa(clock.Hour), strconv.Itoa(clock.Minute)

  if clock.Hour < 10 {
    hour = "0" + hour
  }

  if clock.Minute < 10 {
    minute = "0" + minute
  }

  return fmt.Sprintf("%s:%s", hour, minute)
}
