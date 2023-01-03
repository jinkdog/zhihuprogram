package user

type Group struct{}

func (g *Group) Sign() *SignApi {
	return &insSign
}

//在同一个文件夹下由sign.go引入SignAPi
