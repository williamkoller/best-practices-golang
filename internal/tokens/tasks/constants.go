package tasks

const TypeUpdateToken = "update:token"

type TaskUpdateTokenPayload struct {
	Tokens []string `json:"tokens"`
}
