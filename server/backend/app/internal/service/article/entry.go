package article

type Group struct{}

func (g *Group) Article() *SArticle {
	return &insArticle
}
