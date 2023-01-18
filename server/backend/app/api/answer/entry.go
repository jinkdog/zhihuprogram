package answer

type Group struct{}

func (g *Group) Show() *ShowApi {
	return &insShow
}
func (g *Group) Change() *ChangeApi {
	return &insChange
}
