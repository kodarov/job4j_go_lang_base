package tracker

import "context"

type UI struct {
	In    Input
	Out   Output
	Store Store
}

func (u UI) Run(ctx context.Context) error {
	actions := map[string]Usecase{
		"add": AddUsecase{},
		"get": GetUsecase{},
	}

	for {
		u.Out.Out("select action (add, get, exit):")
		selected := u.In.Get()

		if selected == "exit" {
			break
		}

		action, ok := actions[selected]
		if !ok {
			u.Out.Out("not found action")
			continue
		}

		if err := action.Done(ctx, u.In, u.Out, u.Store); err != nil {
			return err
		}
	}

	return nil
}
