package github

type Sorter struct {
	SortParam string
	Repos     []Repo
}

func (s Sorter) Len() int {
	return len(s.Repos)
}

func (s Sorter) Less(i, j int) bool {
	switch s.SortParam {
	case "stars":
		return s.Repos[i].Stars > s.Repos[j].Stars
	case "forks":
		return s.Repos[i].Forks > s.Repos[j].Forks
	case "watchers":
		return s.Repos[i].Watchers > s.Repos[j].Watchers
	case "issues":
		return s.Repos[i].Issues > s.Repos[j].Issues
	}

	return false
}

func (s Sorter) Swap(i, j int) {
	s.Repos[i], s.Repos[j] = s.Repos[j], s.Repos[i]
}
