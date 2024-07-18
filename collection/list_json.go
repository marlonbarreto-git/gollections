package collection

import "encoding/json"

func (list *List[T]) UnmarshalJSON(bytes []byte) (err error) {
	var target []T

	if err = json.Unmarshal(bytes, &target); err != nil {
		return
	}

	*list = target
	return
}

func (list *List[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(list.ToArray())
}
