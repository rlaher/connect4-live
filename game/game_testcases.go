package game

var testcases = []struct {
	mygame        [BoardHeight][BoardWidth]Slot
	playersymbols []string
	result        bool
	pos           []int
	player        int
}{
	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{0, 0},
		1,
	},

	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},

		[]string{0: "X", 1: "O"},

		true,
		[]int{0, 0},
		1,
	},
	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{0, 0},
		1,
	},
	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		false,
		[]int{5, 4},
		0,
	},
	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "X"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{3, 3},
		1,
	},
	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{0, 3},
		1,
	},
	{
		[BoardHeight][BoardWidth]Slot{
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "X"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
			{{true, "O"}, {true, "X"}, {true, "X"}, {true, "O"}, {true, "O"}, {true, "O"}, {true, "X"}},
		},
		[]string{0: "X", 1: "O"},
		true,
		[]int{6, 0},
		0,
	},
}
