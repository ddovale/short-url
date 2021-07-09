package urls

type Url struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

type UrlsList []Url

var (
	urls       = UrlsList{}
	lastId int = 0
)

func All() UrlsList {
	return urls
}

func Get(id int) *Url {
	for _, url := range urls {
		if url.ID == id {
			return &url
		}
	}

	return nil
}

func Create(url *Url) {
	lastId = lastId + 1
	url.ID = lastId
	urls = append(urls, *url)
}
