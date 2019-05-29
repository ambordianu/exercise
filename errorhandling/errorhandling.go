package errorhandling

import "encoding/json"

type ErrOrderNotFound struct {
	error
}

func OrderNotFoundError(e error) error {
	return ErrOrderNotFound{e}
}

func (e ErrOrderNotFound) MarshalJSON() ([]byte, error) {
	output := map[string]string{}
	output["error"] = e.Error()
	return json.Marshal(output)
}
