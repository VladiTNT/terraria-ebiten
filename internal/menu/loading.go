package menu

type LoadingAnimation struct {
	FrameCounter int
	Frames       []string
}

func NewLoadingAnimation() LoadingAnimation {
	return LoadingAnimation{
		FrameCounter: 0,
		Frames: []string{
			"Waiting",
			"Waiting.",
			"Waiting..",
			"Waiting...",
		},
	}
}

func (ls *LoadingAnimation) Update() error {
	ls.FrameCounter++

	if ls.FrameCounter >= 120 {
		ls.FrameCounter = 0
	}

	return nil
}
