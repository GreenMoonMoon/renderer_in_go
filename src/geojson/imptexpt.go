package geojson

import (
    "encoding/json"
    "errors"
)

func Export(i interface{}) (string, error) {
    if i == nil {
        return "", errors.New("Parameter is nil")
    }
    a, b := json.Marshal(i)
    return string(a), b
}