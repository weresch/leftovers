package fakes

type StaticRoutes struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			RouterID string
		}
		Returns struct {
			Error error
		}
	}
}

func (s *StaticRoutes) Delete(routerID string) error {
	s.DeleteCall.CallCount++

	s.DeleteCall.Receives.RouterID = routerID
	return s.DeleteCall.Returns.Error
}
