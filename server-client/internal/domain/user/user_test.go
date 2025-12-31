package user

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {
	validUserID := uuid.New()

	tests := []struct {
		name             string
		userID           uuid.UUID
		username         string
		imgurl           string
		pronunciation    string
		selfintroduction string
		stars            float32
		wantErr          string
	}{
		{
			name:             "正常系",
			userID:           validUserID,
			username:         "John Doe",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "暇人です",
			selfintroduction: "こんにちは、ジョンです",
			stars:            50,
			wantErr:          "",
		},
		{
			name:             "UserIDがNil",
			userID:           uuid.Nil,
			username:         "John Doe",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "暇人です",
			selfintroduction: "こんにちは、ジョンです",
			stars:            50,
			wantErr:          "ユーザーIDが空になっています",
		},
		{
			name:             "Usernameが空",
			userID:           validUserID,
			username:         "",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "暇人です",
			selfintroduction: "こんにちは、ジョンです",
			stars:            50,
			wantErr:          "ユーザー名が空になっています",
		},
		{
			name:             "Imgurlが空",
			userID:           validUserID,
			username:         "John Doe",
			imgurl:           "",
			pronunciation:    "暇人です",
			selfintroduction: "こんにちは、ジョンです",
			stars:            50,
			wantErr:          "画像URLが空になっています",
		},
		{
			name:             "Pronunciationが空",
			userID:           validUserID,
			username:         "John Doe",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "",
			selfintroduction: "こんにちは、ジョンです",
			stars:            50,
			wantErr:          "代名詞が空になっています",
		},
		{
			name:             "Selfintroductionが空",
			userID:           validUserID,
			username:         "John Doe",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "暇人です",
			selfintroduction: "",
			stars:            50,
			wantErr:          "自己紹介が空になっています",
		},
		{
			name:             "Starsが範囲外（大きすぎ）",
			userID:           validUserID,
			username:         "John Doe",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "暇人です",
			selfintroduction: "こんにちは、ジョンです",
			stars:            150,
			wantErr:          "評価が無効です",
		},
		{
			name:             "Starsが範囲外（小さすぎ）",
			userID:           validUserID,
			username:         "John Doe",
			imgurl:           "https://example.com/img.jpg",
			pronunciation:    "暇人です",
			selfintroduction: "こんにちは、ジョンです",
			stars:            -150,
			wantErr:          "評価が無効です",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser(tt.userID, tt.username, tt.imgurl, tt.pronunciation, tt.selfintroduction, tt.stars)

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
			if u.Userid() != tt.userID {
				t.Errorf("%s: expected UserID %v, got %v", tt.name, tt.userID, u.Userid())
			}
			if u.Username() != tt.username {
				t.Errorf("%s: expected Username %q, got %q", tt.name, tt.username, u.Username())
			}
			if u.Imgurl() != tt.imgurl {
				t.Errorf("%s: expected Imgurl %q, got %q", tt.name, tt.imgurl, u.Imgurl())
			}
			if u.Pronunciation() != tt.pronunciation {
				t.Errorf("%s: expected Pronunciation %q, got %q", tt.name, tt.pronunciation, u.Pronunciation())
			}
			if u.Selfintroduction() != tt.selfintroduction {
				t.Errorf("%s: expected Selfintroduction %q, got %q", tt.name, tt.selfintroduction, u.Selfintroduction())
			}
			if u.Stars() != tt.stars {
				t.Errorf("%s: expected Stars %v, got %v", tt.name, tt.stars, u.Stars())
			}
		})
	}
}
