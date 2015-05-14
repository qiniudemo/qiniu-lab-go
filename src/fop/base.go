package fop

const (
	PFOP_URL = "http://api.qiniu.com/pfop/"
)

type PfopResult struct {
	PersistentId string `json:"persistentId"`
}
