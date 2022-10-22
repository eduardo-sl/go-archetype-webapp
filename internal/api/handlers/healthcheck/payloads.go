package healthcheck

type status string

const (
	pass = "pass"
	fail = "fail"
)

type response struct {
	Status status `json:"status"`
}