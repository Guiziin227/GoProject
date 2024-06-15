package modelos

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //omitempty se vazio nao passa para o json.
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}
