package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"isdc.fi/splitbit/server/api"
	"isdc.fi/splitbit/server/data"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target api --clean ../openapi.yml

var qs *data.Queries

type Handler struct {
	api.UnimplementedHandler // automatically implement all methods by embedding a struct
}

func (h *Handler) GetPing(ctx context.Context) (*api.PongResponse, error) {
	return &api.PongResponse{Message: "value"}, nil
}

func (h *Handler) GroupsGet(ctx context.Context) ([]api.Group, error) {
	g, _ := qs.GetGroupsOfMember(ctx, ctx.Value("user_id").(int64))

	resp := make([]api.Group, len(g))
	// Convert each database Group to API Group
	for i, dbGroup := range g {

		params := data.GetNetAmountForUserInGroupParams{
			GroupID:  dbGroup.ID,
			AuthorID: ctx.Value("user_id").(int64),
		}

		amountOwed, err := qs.GetNetAmountForUserInGroup(ctx, params)
		if err != nil {
			return nil, err
		}

		resp[i] = api.Group{
			ID:         int(dbGroup.ID),
			Name:       dbGroup.Name,
			NoItems:    !amountOwed.Valid,
			AmountOwed: amountOwed.Float64,
		}
	}
	return resp, nil
}

func (h *Handler) GroupsPost(ctx context.Context, req *api.GroupsPostReq) (int, error) {
	gid, _ := qs.AddGroup(ctx, req.Name)

	return int(gid), nil
}

func (h *Handler) GroupsNonauthedGet(ctx context.Context) ([]api.Group, error) {
	g, _ := qs.GetGroupsOfMember(ctx, 1)
	resp := make([]api.Group, len(g))
	// Convert each database Group to API Group
	for i, v := range g {
		resp[i] = api.Group{
			ID:   int(v.ID),
			Name: v.Name,
		}
	}
	return resp, nil
}

func (h *Handler) GroupsIDGet(ctx context.Context, params api.GroupsIDGetParams) (*api.GroupOverview, error) {
	//IDs
	g_id := int64(params.ID)
	userid := ctx.Value("user_id").(int64)

	//authorID
	group, _ := qs.GetGroupByID(ctx, g_id)

	// Fetch members of the group
	members, _ := qs.GetMembersOfGroup(ctx, g_id)

	var apiMembers []api.Member

	for _, member := range members {
		apiMembers = append(apiMembers, api.Member{
			ID:   int(member.ID),
			Name: member.Username,
		})
	}
	// new schema EP added converting to of type memeber th query of Expense_participants
	eps, _ := qs.Expense_participants(ctx, 1)
	var ep_members []api.Member

	for _, ep := range eps {
		ep_members = append(ep_members, api.Member{
			ID:          int(ep.MemberID),
			Name:        ep.Username,
			DisplayName: ep.Displayname,
		})

	}
	// Fetch items of the group
	items, _ := qs.GetItemsOfGroup(ctx, g_id)

	var apiItems []api.Item

	for _, item := range items {
		apiItems = append(apiItems, api.Item{
			ID:        int(item.ID),
			Timestamp: int(item.Timestamp),
			Name:      item.Name,
			Price:     item.Price,
			AuthorID:  int(item.AuthorID),
			GroupID:   int(g_id),
			// The new participant list!!!!!!
			Participants: ep_members,
		})
	}

	// Calculate money balance for the group
	netAmountParams := data.GetNetAmountForUserInGroupParams{
		AuthorID: userid,
		GroupID:  g_id,
	}

	netAmount, err := qs.GetNetAmountForUserInGroup(ctx, netAmountParams)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch net amount: %w", err)
	}
	var balance int

	if netAmount.Valid {
		balance = int(netAmount.Float64)
	} else {
		balance = 0
	}

	groupOverview := &api.GroupOverview{
		Name:         group.Name,
		Members:      apiMembers,
		Items:        apiItems,
		MoneyBalance: balance,
	}

	return groupOverview, nil
}

func (h *Handler) GroupsIDMembersPost(ctx context.Context, req *api.GroupsIDMembersPostReq, params api.GroupsIDMembersPostParams) error {
	qs.AddMemberToGroup(ctx, data.AddMemberToGroupParams{GroupID: int64(params.ID), MemberID: int64(req.MemberID)})
	return nil
}

func (h *Handler) GroupsIDItemsGet(ctx context.Context, params api.GroupsIDItemsGetParams) ([]api.Item, error) {
	g, _ := qs.GetItemsOfGroup(ctx, int64(params.ID))

	resp := make([]api.Item, len(g))
	for i, v := range g {
		resp[i] = api.Item{
			ID:        int(v.ID),
			Timestamp: int(v.Timestamp),
			Name:      v.Name,
			Price:     v.Price,
			AuthorID:  int(v.AuthorID),
			GroupID:   int(v.GroupID),
		}
	}
	return resp, nil
}

func (h *Handler) GroupsIDItemsPost(ctx context.Context, req *api.Item, params api.GroupsIDItemsPostParams) (int, error) {
	g, _ := qs.AddItemToGroup(ctx, data.AddItemToGroupParams{
		Name:          req.Name,
		Timestamp:     int64(req.Timestamp),
		Price:         req.Price,
		GroupID:       int64(params.ID),
		AuthorID:      int64(req.AuthorID),
		Reimbursement: sql.NullBool{Bool: req.Reimbursement.Value, Valid: req.Reimbursement.Set},
	})

	return int(g), nil
}

func main() {
	// Create service instance.

	db := setupDB()

	qs = data.New(db) // qs stands for queries
	service := &Handler{}
	sec := &Security{}
	// Create generated server.
	srv, err := api.NewServer(service, sec)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("SPLITBIT_PORT")
	if port == "" {
		port = ":8080"
	}
	log.Printf("\033[32mSplitBit server has started on port %s\033[m\n", port)
	if err := http.ListenAndServe(port, corsMiddleware(srv)); err != nil {
		log.Fatal(err)
	}
}
