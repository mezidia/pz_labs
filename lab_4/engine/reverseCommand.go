package engine

type ReverseCommand struct {
	Arg string
}

func (p *ReverseCommand) Execute(loop Handler) {
	runes := []rune(p.Arg)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	var print PrintCommand
	print.Arg = string(runes)
	print.Execute(loop)
}
