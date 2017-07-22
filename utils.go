package main

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

func makeInterface(fields []string) []interface{} {
	vals := make([]interface{}, len(fields))
	for i, v := range fields {
		vals[i] = v
	}
	return vals
}

func parsePublishedAt(ts string) string {
	var pubAt string
	if ts == "" {
		t := time.Now()
		p := t.Format(time.RFC3339)
		pubAt = strings.Replace(p, "T", " ", 1)
		pubAt = pubAt[:19] + ".555555"
	} else {
		pubAt = strings.Replace(ts, "T", " ", 1)
		pubAt = pubAt[:19] + ".555555"
	}
	return pubAt
}

func int64ToStr(i int64) string {
	x := strconv.FormatInt(i, 10)
	return x
}
func rQ(s string) string {
	// need to replace ' with '' before inserting to db
	r := strings.NewReplacer(`&#039;`, `''`, `\u00a0`, ` `, `'`, `''`)
	result := r.Replace(s)
	return result
}
func rU(s string) string {
	replaced := strings.Replace(s, `\n\thttp`, `http`, -1)
	u, _ := url.Parse(replaced)
	result := strings.Replace(replaced, u.RawQuery, ``, -1)
	return result
}
