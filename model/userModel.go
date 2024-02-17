package model

type UserModel struct {
	UserName string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}
