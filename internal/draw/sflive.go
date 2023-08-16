package draw

import "time"

// Set4Life represents a draw from Set for Life
type Set4Life struct {
	DrawDate  time.Time    `json:"draw_date" sqlite:"draw_date,INTEGER"`
	DayOfWeek time.Weekday `json:"day_of_week" sqlite:"day_of_week,INTEGER"`
	Ball1     uint8        `json:"ball1" sqlite:"ball1,INTEGER"`
	Ball2     uint8        `json:"ball2" sqlite:"ball2,INTEGER"`
	Ball3     uint8        `json:"ball3" sqlite:"ball3,INTEGER"`
	Ball4     uint8        `json:"ball4" sqlite:"ball4,INTEGER"`
	Ball5     uint8        `json:"ball5" sqlite:"ball5,INTEGER"`
	LifeBall  uint8        `json:"life_ball" sqlite:"life_ball,INTEGER"`
	BallSet   string       `json:"ball_set" sqlite:"ball_set,TEXT"`
	Machine   string       `json:"machine" sqlite:"machine,TEXT"`
	DrawNo    uint64       `json:"draw_no" sqlite:"draw_no,TEXT"`
}
