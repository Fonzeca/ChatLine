package model

import "time"

type UserIds struct {
	Ids []string `json:"ids"`
}

type ComentarioMQ struct {
	UsuarioID int32     `json:"usuario_id"`
	Fecha     time.Time `json:"fecha"`
	Tema      string    `json:"tema"`
	TemaId    int32     `json:"tema_id"`
	Mensaje   string    `json:"mensaje"`
}
