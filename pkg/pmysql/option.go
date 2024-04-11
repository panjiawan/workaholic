package pmysql

func WithConnection(host string, port int, username, password, db string) Option {
	return func(s *Service) {
		s.host = host
		s.port = port
		s.username = username
		s.password = password
		s.db = db
	}
}

func WithPrefix(prefix string) Option {
	return func(s *Service) {
		s.prefix = prefix
	}
}

func WithLimit(maxIdle, maxOpen int) Option {
	return func(s *Service) {
		s.maxIdle = maxIdle
		s.maxOpen = maxOpen
	}
}

func WithCharset(charset string) Option {
	return func(s *Service) {
		s.charset = charset
	}
}

func WithDebug(debug bool) Option {
	return func(s *Service) {
		s.debug = debug
	}
}
