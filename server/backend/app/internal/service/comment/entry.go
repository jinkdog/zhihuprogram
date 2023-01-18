package comment

type Group struct{}

func (g *Group) Comment() *SComment {
	return &insComment
}
