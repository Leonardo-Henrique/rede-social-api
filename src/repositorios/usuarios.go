package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	stmt, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer stmt.Close()

	resultado, erro := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDinserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDinserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criado_em FROM usuarios WHERE nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {

		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil

}

func (repositorio Usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT  id, nome, nick, email, criado_em FROM usuarios WHERE id= ?", id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) AtualizarUsuario(id uint64, usuario modelos.Usuario) error {
	stmt, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, id)
	if erro != nil {
		return erro
	}

	return nil

}

func (repositorio Usuarios) DeletarUsuario(id uint64) error {
	stmt, erro := repositorio.db.Prepare(
		"DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID uint64, seguidorID uint64) error {
	stmt, erro := repositorio.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES(?,?)")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil

}

func (repositorio Usuarios) DeixarDeSeguir(usuarioID uint64, seguidorID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		"DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?")
	if erro != nil {
		return nil
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email FROM usuarios u 
		INNER JOIN seguidores s on u.id = s.seguidor_id 
		WHERE s.usuario_id =?
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario

	for linhas.Next() {
		var seguidor modelos.Usuario
		if linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
		); erro != nil {
			return nil, erro
		}

		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil

}

func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email 
		FROM usuarios u INNER JOIN seguidores s
		ON u.id = s.usuario_id WHERE s.seguidor_id = ? 
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarSenhaPorID(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query(
		`SELECT senha FROM usuarios WHERE id = ?`, usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil

}

func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, novaSenha string) error {
	stmt, erro := repositorio.db.Prepare(`
		UPDATE usuarios SET senha = ? WHERE id = ?
	`)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(novaSenha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
