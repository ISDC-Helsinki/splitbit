// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package data

type Group struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID        int64   `json:"id"`
	Timestamp int64   `json:"timestamp"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	AuthorID  int64   `json:"author_id"`
	GroupID   int64   `json:"group_id"`
}

type Member struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
	Password    string `json:"password"`
}

type MemberGroup struct {
	GroupID  int64 `json:"group_id"`
	MemberID int64 `json:"member_id"`
}