import "sync"

var _instance *service
var once sync.Once

type service struct {
	Name string
}

func Service() *service {
	once.Do(func() {
		_instance = &service{}
	})
	return _instance
}

func (this *service) Startup() {
}

func (this *service) Stop() {

}

func (this *service) Restart() {

}
