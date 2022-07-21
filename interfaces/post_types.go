package interfaces

type Post struct {
	Title        string  `json:"title,omitempty" gorm:"column:title"`
	Date         string  `json:"date,omitempty" gorm:"column:date"`
	LastMod      string  `json:"lastmod,omitempty" gorm:"column:lastmod"`
	Tags         []uint8 `json:"tags,omitempty" gorm:"column:tags"`
	Draft        bool    `json:"draft,omitempty" gorm:"column:draft"`
	Summary      string  `json:"summary,omitempty" gorm:"column:summary"`
	Overview     string  `json:"overview" gorm:"column:overview"`
	Layout       string  `json:"layout,omitempty" gorm:"column:layout"`
	Bibliography string  `json:"bibliography,omitempty" gorm:"column:bibliography"`
	CanonicalURL string  `json:"canonicalURL,omitempty" gorm:"column:canonicalUrl"`
	Slug         string  `json:"slug,omitempty" gorm:"column:slug"`
}
