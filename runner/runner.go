package runner

import "os/exec"

type Runner struct {
	Path string
	Params []string
	Stdout []byte
}

func (r *Runner) Run() error {
	var err error

	command := exec.Command(r.Path, r.Params...)

	r.Stdout, err = command.Output()
	if err != nil {
		return err
	}

	return nil
}
