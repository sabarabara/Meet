package recruitlocation_test

import (
	recruitlocation "server-client/internal/domain/vo_and_entity/recruit/recruit_location"
	"testing"

	"github.com/google/uuid"
)

func TestNewRecruitLocation(t *testing.T) {

	tests := []struct {
		name              string
		recruitLocationID uuid.UUID
		recruitID         uuid.UUID
		latitude          float64
		longitude         float64
		wantErr           string
	}{
		{
			name:              "正常系",
			recruitLocationID: uuid.New(),
			recruitID:         uuid.New(),
			latitude:          35.681236,
			longitude:         139.767125,
			wantErr:           "",
		},
		{
			name:              "RecruitLocationIDがNil",
			recruitLocationID: uuid.Nil,
			recruitID:         uuid.New(),
			latitude:          35.681236,
			longitude:         139.767125,
			wantErr:           "募集位置IDが空になっています",
		},
		{
			name:              "RecruitIDがNil",
			recruitLocationID: uuid.New(),
			recruitID:         uuid.Nil,
			latitude:          35.681236,
			longitude:         139.767125,
			wantErr:           "募集IDが空になっています",
		},
		{
			name:              "Latitudeが0",
			recruitLocationID: uuid.New(),
			recruitID:         uuid.New(),
			latitude:          0,
			longitude:         139.767125,
			wantErr:           "緯度が空になっています",
		},
		{
			name:              "Longitudeが0",
			recruitLocationID: uuid.New(),
			recruitID:         uuid.New(),
			latitude:          35.681236,
			longitude:         0,
			wantErr:           "経度が空になっています",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc, err := recruitlocation.NewRecruitLocation(tt.recruitLocationID, tt.recruitID, tt.latitude, tt.longitude)
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
			if loc.Recruitlocationid() != tt.recruitLocationID {
				t.Errorf("%s: expected Recruitlocationid %v, got %v", tt.name, tt.recruitLocationID, loc.Recruitlocationid())
			}
			if loc.Recruitid() != tt.recruitID {
				t.Errorf("%s: expected Recruitid %v, got %v", tt.name, tt.recruitID, loc.Recruitid())
			}
			if loc.Latitude() != tt.latitude {
				t.Errorf("%s: expected Latitude %v, got %v", tt.name, tt.latitude, loc.Latitude())
			}
			if loc.Longitude() != tt.longitude {
				t.Errorf("%s: expected Longitude %v, got %v", tt.name, tt.longitude, loc.Longitude())
			}
		})
	}
}
