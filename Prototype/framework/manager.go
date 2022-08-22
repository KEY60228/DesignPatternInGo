package framework

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{showcase: make(map[string]Product)}
}

func (m *Manager) Register(name string, prototype Product) {
	m.showcase[name] = prototype
}

func (m *Manager) Create(prototypeName string) Product {
	p, exists := m.showcase[prototypeName]
	if !exists {
		return nil
	}
	return p.CreateCopy()
}
