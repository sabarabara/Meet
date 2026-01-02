package review

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewReview(t *testing.T) {
	tests := []struct {
		name            string
		reviewID        uuid.UUID
		roomID          uuid.UUID
		writerID        uuid.UUID
		evaluatedUserID uuid.UUID
		comment         string
		wantErr         string
	}{
		{
			name:            "正常系",
			reviewID:        uuid.New(),
			roomID:          uuid.New(),
			writerID:        uuid.New(),
			evaluatedUserID: uuid.New(),
			comment:         "良い部屋でした！",
			wantErr:         "",
		},
		{
			name:            "ReviewIDがNil",
			reviewID:        uuid.Nil,
			roomID:          uuid.New(),
			writerID:        uuid.New(),
			evaluatedUserID: uuid.New(),
			comment:         "コメント",
			wantErr:         "レビューIDが空になっています",
		},
		{
			name:            "RoomIDがNil",
			reviewID:        uuid.New(),
			roomID:          uuid.Nil,
			writerID:        uuid.New(),
			evaluatedUserID: uuid.New(),
			comment:         "コメント",
			wantErr:         "ルームIDが空になっています",
		},
		{
			name:            "WriterIDがNil",
			reviewID:        uuid.New(),
			roomID:          uuid.New(),
			writerID:        uuid.Nil,
			evaluatedUserID: uuid.New(),
			comment:         "コメント",
			wantErr:         "書き込み者IDが空になっています",
		},
		{
			name:            "EvaluatedUserIDがNil",
			reviewID:        uuid.New(),
			roomID:          uuid.New(),
			writerID:        uuid.New(),
			evaluatedUserID: uuid.Nil,
			comment:         "コメント",
			wantErr:         "評価されるユーザーIDが空になっています",
		},
		{
			name:            "コメントが空",
			reviewID:        uuid.New(),
			roomID:          uuid.New(),
			writerID:        uuid.New(),
			evaluatedUserID: uuid.New(),
			comment:         "",
			wantErr:         "コメントが空になっています",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewReview(tt.reviewID, tt.roomID, tt.writerID, tt.evaluatedUserID, tt.comment)

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
			if r.Reviewid() != tt.reviewID {
				t.Errorf("%s: expected ReviewID %v, got %v", tt.name, tt.reviewID, r.Reviewid())
			}
			if r.Roomid() != tt.roomID {
				t.Errorf("%s: expected RoomID %v, got %v", tt.name, tt.roomID, r.Roomid())
			}
			if r.Writer() != tt.writerID {
				t.Errorf("%s: expected WriterID %v, got %v", tt.name, tt.writerID, r.Writer())
			}
			if r.EvaluatedUser() != tt.evaluatedUserID {
				t.Errorf("%s: expected EvaluatedUserID %v, got %v", tt.name, tt.evaluatedUserID, r.EvaluatedUser())
			}
			if r.Comment() != tt.comment {
				t.Errorf("%s: expected Comment %q, got %q", tt.name, tt.comment, r.Comment())
			}
		})
	}
}
