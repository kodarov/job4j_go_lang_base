package tracker

type UI struct {
	In      Input
	Out     Output
	Tracker *Tracker
}

func (u UI) Run() {
	actions := map[string]Usecase{
		"add":    AddUsecase{},
		"get":    GetUsecase{},
		"update": UpdateUsecase{},
		"delete": DeleteUsecase{},
		"del":    DeleteUsecase{},
		"search": SearchUsecase{},
		"find":   SearchUsecase{},
	}

	for {
		u.Out.Out("select action")
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("not found action")
			continue
		}

		action.Done(u.In, u.Out, u.Tracker)
	}
}
