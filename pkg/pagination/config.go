package pagination

const DefaultKeyPage = "page"
const DefaultKeyPerPage = "per_page"
const DefaultPerPage = 100

type Config struct {
	KeyPage    string
	KeyPerPage string
	Page       uint
	PerPage    uint
}

func (c Config) keyPage() string {
	if c.KeyPage == "" {
		return DefaultKeyPage
	}

	return c.KeyPage
}

func (c Config) keyPerPage() string {
	if c.KeyPerPage == "" {
		return DefaultKeyPerPage
	}

	return c.KeyPerPage
}

func (c Config) page() uint {
	if c.Page == 0 {
		return 1
	}

	return c.Page
}

func (c Config) perPage() uint {
	if c.PerPage == 0 {
		return DefaultPerPage
	}
	return c.PerPage
}
