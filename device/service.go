package device

// ServiceFn process packet and return serivce id and drop flag
type ServiceFn func(buff []byte) (service uint64, shouldDrop bool)

var serviceFns []ServiceFn

// RegisterServiceFn register service function to identify packet
func RegisterServiceFn(fn ServiceFn) {
	serviceFns = append(serviceFns, fn)
}

// ExecuteServiceFns to process packet data
func ExecuteServiceFns(buff []byte) (service uint64, shouldDrop bool) {
	finalService := uint64(0)
	for _, fn := range serviceFns {
		service, shouldDrop = fn(buff)
		if service != 0 {
			finalService = service
		}
		if shouldDrop {
			return finalService, true
		}
	}
	return finalService, false
}
