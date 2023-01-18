package answer_post

type Group struct{}

func (g *Group) Answer() *SAnswer {
	return &insAnswer
}
