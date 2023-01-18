package question_community

type Group struct{}

func (g *Group) Question() *SQuestion {
	return &insQuestion
}
