package factory

type IFactory interface {
	CreateLink(string, string) ILink
	CreateTray(string) ITray
	CreatePage(string, string) IPage
}
