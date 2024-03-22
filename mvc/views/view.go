package views

type View interface {
	Render() error
	template() string
}
