
# Rede Social - Guia de referência da API

Uma breve descrição do que este projeto faz e para quem se destina


## Referência da API

#### Criar um novo usuário

```http
  POST /usuarios
```

*Body request:*
```json
{
	"nome": "your name",
	"nick": "your nick",
	"email": "your@email.com",
	"senha": "your_pass"
}
```



#### Seguir um usuário

```http
  POST /usuarios/${usuarioId}/seguir
```

| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 


#### Parar de seguir um usuário

```http
  POST /usuarios/${usuarioId}/parar-de-seguir
```

| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 
  
  #### Atualizar senha

```http
  POST /usuarios/${usuarioId}/atualizar-senha
```
*Body request:*
```json
{
	"senha_atual": "senha_atual",
	"nova_senha": "nova_senha"
}
```

| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 

 #### Listar usuários da rede social por ID

```http
  GET /usuarios/${usuarioId}
```


| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 


 #### Listar todos os usuários da rede social

```http
  GET /usuarios
```
*Body request:*
```json
{
	"senha_atual": "senha_atual",
	"nova_senha": "nova_senha"
}
```

| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 

 #### Listar seguidores de um usuário 

```http
  GET /usuarios/${usuarioId}/seguidores
```
*Body request:*
```json
{
	"senha_atual": "senha_atual",
	"nova_senha": "nova_senha"
}
```

| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 

 #### Listar quem o usuário segue

```http
  GET /usuarios/${usuarioId}/seguindo
```


| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 

#### Atualizar informações do usuário

```http
  PUT /usuarios/${usuarioId}
```


| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 

#### Deletar informações do usuário

```http
  DELETE /usuarios/${usuarioId}
```


| Precisa de autenticação? | Tipo     
| :-------- | :------- |
| **sim**      | `bearer token` | 
