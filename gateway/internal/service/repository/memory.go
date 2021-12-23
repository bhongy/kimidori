package repository

// Memory implements Repository interface using in-memory data structure
// the primary use case is for testing
type Memory struct {
	// service name -> Backend
	serviceMap map[string]Backend
}

func NewMemory(serviceMap map[string]Backend) Memory {
	return Memory{serviceMap: serviceMap}
}

func (m Memory) ByServiceName(svc string) (Backend, error) {
	be, ok := m.serviceMap[svc]
	if !ok {
		return Backend{}, ErrNotFound
	}
	return be, nil
}
