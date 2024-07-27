package user_test

import (
	"fmt"
	"sort"
	"testing"
	"time"
	"wb/user"
)

func TestUserList_Sort(t *testing.T) {
	users := user.UserList{
		{ID: "1", CreatedAt: time.Now().Add(-10 * time.Hour)},
		{ID: "2", CreatedAt: time.Now().Add(-5 * time.Hour)},
		{ID: "3", CreatedAt: time.Now().Add(-1 * time.Hour)},
		{ID: "4", CreatedAt: time.Now().Add(-11 * time.Hour)},
	}

	expected := []string{"4", "1", "2", "3"}

	sort.Sort(users)

	for _, v := range users {
		fmt.Println(v.ID, v.CreatedAt)
	}

	for i, user := range users {
		if user.ID != expected[i] {
			t.Errorf("Expected user with ID %s, but got %s", expected[i], user.ID)
		}
	}
}
