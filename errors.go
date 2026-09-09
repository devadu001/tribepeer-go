package tribepeer

import "fmt"

type Error struct {
	Status  int
	Message string
	Body    map[string]any
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("TribePeer request failed (%d)", e.Status)
}

type PaymentRequiredError struct {
	Status  int
	Message string
	Body    map[string]any
	Upgrade map[string]any
}

func (e *PaymentRequiredError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "TribePeer request failed (402)"
}
