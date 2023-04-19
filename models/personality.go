package models

//Struct que contém o nome e a história de cada pesronalidade que será
//inserida em nossa API. E indicamos que iremos enviar esses parâmetros
//por meio de um arquivo json
type Personality struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

//Variável do tipo array Personality, que é nossa struct.
var Personalidades []Personality
