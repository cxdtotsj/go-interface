package run

import (
	simplejson "github.com/bitly/go-simplejson"
)

// GetRespField range the response by the key
func GetRespField(j *simplejson.Json, keys []string) (actuals []*simplejson.Json) {
	for _, k := range keys {
		v := j.Get(k)
		actuals = append(actuals, v)
	}
	return actuals
}
