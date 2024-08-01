package main

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05Z07:00"))
	return []byte(stamp), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	var err error
	now, err := time.Parse("2006-01-02T15:04:05Z07:00", string(data))
	*t = JsonTime(now)
	return err
}

func main() {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	fmt.Println(tomorrow)
	fmt.Println(now.After(tomorrow))

	t2 := time.Now().Add(time.Minute * 2)
	fmt.Println(t2.After(now))
}
