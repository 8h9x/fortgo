package fortgo

import "github.com/8h9x/fortgo/friends"

type ClientUser struct {
	client *Client
	id     string
}

func (u *ClientUser) ID() string {
	return u.id
}

func (u *ClientUser) FetchFriends() ([]friends.Friend, error) {
	return u.client.FriendService.GetFriendsList(u.ID())
}
