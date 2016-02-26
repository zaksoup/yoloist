package runner

type Runner struct {
	Path string
	Params []string
	Stdout []byte
}

func (r Runner) Run() error {
	return nil
}
