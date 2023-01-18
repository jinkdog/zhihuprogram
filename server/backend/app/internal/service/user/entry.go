package user

type Group struct{}

func (g *Group) User() *SUser {
	return &insUser
}
func (g *Group) Collect() *SCollect {
	return &insCollect
}
