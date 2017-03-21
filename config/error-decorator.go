package config

type HttpErrorHandler interface {
	handle(p func()) (interface{}, error)
}

type DefaultHandler struct {
}

func (handler* DefaultHandler) Handle(p func()) (interface{}, error) {
	 a, err := p();
	 if(err != nil) {
		  return a, nil;
	 } else {
			return nil, 404;
	 }
}
