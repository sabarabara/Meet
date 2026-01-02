package room

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewRoom(t *testing.T) {
	tests := []struct {
		name       string
		roomID     uuid.UUID
		recruitID  uuid.UUID
		userID     uuid.UUID
		role       string
		isFinished bool
		wantErr    string
	}{
		{
			name:       "正常系",
			roomID:     uuid.New(),
			recruitID:  uuid.New(),
			userID:     uuid.New(),
			role:       "host",
			isFinished: false,
			wantErr:    "",
		},
		{
			name:       "RoomIDがNil",
			roomID:     uuid.Nil,
			recruitID:  uuid.New(),
			userID:     uuid.New(),
			role:       "host",
			isFinished: false,
			wantErr:    "room ID cannot be empty",
		},
		{
			name:       "RecruitIDがNil",
			roomID:     uuid.New(),
			recruitID:  uuid.Nil,
			userID:     uuid.New(),
			role:       "host",
			isFinished: false,
			wantErr:    "recruit ID cannot be empty",
		},
		{
			name:       "UserIDがNil",
			roomID:     uuid.New(),
			recruitID:  uuid.New(),
			userID:     uuid.Nil,
			role:       "host",
			isFinished: false,
			wantErr:    "user ID cannot be empty",
		},
		{
			name:       "Roleが空",
			roomID:     uuid.New(),
			recruitID:  uuid.New(),
			userID:     uuid.New(),
			role:       "",
			isFinished: false,
			wantErr:    "role cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewRoom(tt.roomID, tt.recruitID, tt.userID, tt.role, tt.isFinished)

			if tt.wantErr != "" {
				if err == nil {
					t.Errorf("%s: expected error %q, got nil", tt.name, tt.wantErr)
				} else if err.Error() != tt.wantErr {
					t.Errorf("%s: expected error %q, got %q", tt.name, tt.wantErr, err.Error())
				} else {
					t.Logf("%s: returned error: %v", tt.name, err)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: expected no error, got %v", tt.name, err)
				return
			}

			// 正常系フィールドチェック
			if r.Roomid() != tt.roomID {
				t.Errorf("%s: expected RoomID %v, got %v", tt.name, tt.roomID, r.Roomid())
			}
			if r.Recruitid() != tt.recruitID {
				t.Errorf("%s: expected RecruitID %v, got %v", tt.name, tt.recruitID, r.Recruitid())
			}
			if r.Userid() != tt.userID {
				t.Errorf("%s: expected UserID %v, got %v", tt.name, tt.userID, r.Userid())
			}
			if r.Role() != tt.role {
				t.Errorf("%s: expected Role %q, got %q", tt.name, tt.role, r.Role())
			}
			if r.Isfinished() != tt.isFinished {
				t.Errorf("%s: expected Isfinished %v, got %v", tt.name, tt.isFinished, r.Isfinished())
			}
		})
	}
}
