package pagination

const DefaultKeyPage = "page"
const DefaultKeyPerPage = "per_page"

type Config struct {
	KeyPage    string
	KeyPerPage string
	Page       int
	PerPage    int
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

func (c Config) page() int {
	if c.Page == 0 {
		return 1
	}

	return c.Page
}

func (c Config) perPage() int {
	return c.PerPage
}
