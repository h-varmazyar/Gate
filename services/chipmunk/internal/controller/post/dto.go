package post

type CreateLinkReq struct {
	Key     string `json:"key"`
	RealUrl string `json:"real_url"`
}

type CreateLinkResp struct {
	Url       string
	Key       string
	Immediate bool
}

type FetchLinkReq struct {
	Key string
}

type Link struct {
	Url       string
	Immediate bool
}
