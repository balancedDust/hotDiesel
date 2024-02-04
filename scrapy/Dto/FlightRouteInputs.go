package dto

type RouteInput struct {
	FromCity string `json:"fromCity"`
	ToCity   string `json:"toCity"`
	FromDate string `json:"fromDate"`
}

type Query struct {
	QueryId    string `json:"queryId"`
	RouteInput `json:"routeInput"`
}
