package recruit

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewRecruit_Validations(t *testing.T) {
	tests := []struct {
		name        string
		recruitID   uuid.UUID
		userID      uuid.UUID
		area        string
		imgurl      string
		man         uint16
		woman       uint16
		vacantMan   uint16
		vacantWoman uint16
		commnt      string
		date        time.Time
		wantErr     string
	}{
		{
			name:        "EmptyRecruitID",
			recruitID:   uuid.Nil,
			userID:      uuid.New(),
			area:        "Tokyo",
			imgurl:      "https://example.com/img.png",
			man:         5,
			woman:       5,
			vacantMan:   2,
			vacantWoman: 3,
			commnt:      "Test",
			date:        time.Now(),
			wantErr:     errorEmptyRecruitID,
		},
		{
			name:        "EmptyUserID",
			recruitID:   uuid.New(),
			userID:      uuid.Nil,
			area:        "Tokyo",
			imgurl:      "https://example.com/img.png",
			man:         5,
			woman:       5,
			vacantMan:   2,
			vacantWoman: 3,
			commnt:      "Test",
			date:        time.Now(),
			wantErr:     errorEmptyUserID,
		},
		{
			name:        "EmptyArea",
			recruitID:   uuid.New(),
			userID:      uuid.New(),
			area:        "",
			imgurl:      "https://example.com/img.png",
			man:         5,
			woman:       5,
			vacantMan:   2,
			vacantWoman: 3,
			commnt:      "Test",
			date:        time.Now(),
			wantErr:     errorEmptyArea,
		},
		{
			name:        "EmptyImgurl",
			recruitID:   uuid.New(),
			userID:      uuid.New(),
			area:        "Tokyo",
			imgurl:      "",
			man:         5,
			woman:       5,
			vacantMan:   2,
			vacantWoman: 3,
			commnt:      "Test",
			date:        time.Now(),
			wantErr:     errorEmptyImgurl,
		},
		{
			name:        "EmptyDate",
			recruitID:   uuid.New(),
			userID:      uuid.New(),
			area:        "Tokyo",
			imgurl:      "https://example.com/img.png",
			man:         5,
			woman:       5,
			vacantMan:   2,
			vacantWoman: 3,
			commnt:      "Test",
			date:        time.Time{}, // ゼロ値
			wantErr:     errorEmptyDate,
		},
		{
			name:        "ValidRecruit",
			recruitID:   uuid.New(),
			userID:      uuid.New(),
			area:        "Tokyo",
			imgurl:      "https://example.com/img.png",
			man:         5,
			woman:       5,
			vacantMan:   2,
			vacantWoman: 3,
			commnt:      "Test",
			date:        time.Now(),
			wantErr:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec, err := NewRecruit(tt.recruitID, tt.userID, tt.area, tt.imgurl, tt.man, tt.woman, tt.vacantMan, tt.vacantWoman, tt.commnt, tt.date)
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

			// 正常系のフィールドチェック
			if rec.Recruitid() != tt.recruitID {
				t.Errorf("%s: expected Recruitid %v, got %v", tt.name, tt.recruitID, rec.Recruitid())
			}
			if rec.Userid() != tt.userID {
				t.Errorf("%s: expected Userid %v, got %v", tt.name, tt.userID, rec.Userid())
			}
			if rec.Area() != tt.area {
				t.Errorf("%s: expected Area %q, got %q", tt.name, tt.area, rec.Area())
			}
			if rec.Imgurl() != tt.imgurl {
				t.Errorf("%s: expected Imgurl %q, got %q", tt.name, tt.imgurl, rec.Imgurl())
			}
			if rec.Man() != tt.man {
				t.Errorf("%s: expected Man %v, got %v", tt.name, tt.man, rec.Man())
			}
			if rec.Woman() != tt.woman {
				t.Errorf("%s: expected Woman %v, got %v", tt.name, tt.woman, rec.Woman())
			}
			if rec.VacantMan() != tt.vacantMan {
				t.Errorf("%s: expected VacantMan %v, got %v", tt.name, tt.vacantMan, rec.VacantMan())
			}
			if rec.VacantWoman() != tt.vacantWoman {
				t.Errorf("%s: expected VacantWoman %v, got %v", tt.name, tt.vacantWoman, rec.VacantWoman())
			}
			if rec.Commnt() != tt.commnt {
				t.Errorf("%s: expected Commnt %q, got %q", tt.name, tt.commnt, rec.Commnt())
			}
			if !rec.Date().Equal(tt.date) {
				t.Errorf("%s: expected Date %v, got %v", tt.name, tt.date, rec.Date())
			}
		})
	}
}
