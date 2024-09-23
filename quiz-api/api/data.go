package api

var NumQuizQuestions = 3

var QuizMock = []Question{
	{
		ID:          "1",
		Description: "Which OS is more popular?",
		Options: []Option{
			{
				ID:          "a",
				Description: "MacOS",
			},
			{
				ID:          "b",
				Description: "Windows",
			},
			{
				ID:          "c",
				Description: "Linux",
			},
		},
	},
	{
		ID:          "2",
		Description: "Which bike was sold more often on the month of July?",
		Options: []Option{
			{
				ID:          "a",
				Description: "Road Bike",
			},
			{
				ID:          "b",
				Description: "Moutain Bike",
			},
		},
	},
	{
		ID:          "3",
		Description: "Which genre buys more apples?",
		Options: []Option{
			{
				ID:          "a",
				Description: "Females",
			},
			{
				ID:          "b",
				Description: "Males",
			},
		},
	},
}

var SolutionMock = map[string]string{
	"1": "a",
	"2": "b",
	"3": "b",
}

var SolutionInit = map[string]string{
	"1": "",
	"2": "",
	"3": "",
}

var UserSolution = SolutionInit

var UserSecretsMock = map[string]string{
	"alice": "rainbow",
	"bob":   "flower",
	"eve":   "boat",
}

var UserAPIKeysMock = map[string]string{
	"alice": "VAFJWEKSFS",
	"bob":   "FEJRGIERGJ",
	"eve":   "PQIENFJRGR",
}

var StatsMock = map[string]float64{
	"alice": 00.00,
	"bob":   00.00,
	"eve":   00.00,
}
