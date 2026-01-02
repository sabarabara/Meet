package message

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewMessage_Validations(t *testing.T) {
	tests := []struct {
		name      string
		messageID uuid.UUID
		roomID    uuid.UUID
		senderID  uuid.UUID
		content   string
		isRead    bool
		wantErr   string
	}{
		{
			name:      "EmptyMessageID",
			messageID: uuid.Nil,
			roomID:    uuid.New(),
			senderID:  uuid.New(),
			content:   "Hello",
			wantErr:   errorEmptyMessageID,
		},
		{
			name:      "EmptyRoomID",
			messageID: uuid.New(),
			roomID:    uuid.Nil,
			senderID:  uuid.New(),
			content:   "Hello",
			wantErr:   errorEmptyRoomID,
		},
		{
			name:      "EmptySenderID",
			messageID: uuid.New(),
			roomID:    uuid.New(),
			senderID:  uuid.Nil,
			content:   "Hello",
			wantErr:   errorEmptySenderID,
		},
		{
			name:      "EmptyContent",
			messageID: uuid.New(),
			roomID:    uuid.New(),
			senderID:  uuid.New(),
			content:   "",
			wantErr:   errorEmptyContent,
		},
		{
			name:      "ValidMessage",
			messageID: uuid.New(),
			roomID:    uuid.New(),
			senderID:  uuid.New(),
			content:   "Hello, World!",
			isRead:    false,
			wantErr:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := NewMessage(tt.messageID, tt.roomID, tt.senderID, tt.content, tt.isRead)
			if tt.wantErr != "" {
				if err == nil {
					t.Errorf("%s: expected error %q, got nil", tt.name, tt.wantErr)
				} else if err.Error() != tt.wantErr {
					t.Errorf("%s: expected error %q, got %q", tt.name, tt.wantErr, err.Error())
				} else {
					// エラー内容も出力して確認したい場合
					t.Logf("%s: returned error: %v", tt.name, err)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: expected no error, got %v", tt.name, err)
				return
			}

			// 正常系のフィールドチェック
			if msg.Messageid() != tt.messageID {
				t.Errorf("%s: expected Messageid %v, got %v", tt.name, tt.messageID, msg.Messageid())
			}
			if msg.Roomid() != tt.roomID {
				t.Errorf("%s: expected Roomid %v, got %v", tt.name, tt.roomID, msg.Roomid())
			}
			if msg.Senderid() != tt.senderID {
				t.Errorf("%s: expected Senderid %v, got %v", tt.name, tt.senderID, msg.Senderid())
			}
			if msg.Content() != tt.content {
				t.Errorf("%s: expected Content %q, got %q", tt.name, tt.content, msg.Content())
			}
			if msg.Isread() != tt.isRead {
				t.Errorf("%s: expected Isread %v, got %v", tt.name, tt.isRead, msg.Isread())
			}
		})
	}
}
