package film

type FilmType string

const (
	NewRelease FilmType = "NEW_RELEASE"
	Regular    FilmType = "REGULAR"
	Old        FilmType = "OLD"
)

type film struct {
	ID    int      `gorm:"primaryKey"`
	Title string   `gorm:"column:name"`
	Type  FilmType `gorm:"column:type"`
}

func (film) TableName() string {
	return "films"
}

func (f *film) dto() FilmDTO {
	return FilmDTO{
		ID:    f.ID,
		Title: f.Title,
		Type:  f.Type,
	}
}
