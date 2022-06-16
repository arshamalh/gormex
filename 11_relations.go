package main

import "fmt"

type StoryStatus string

const (
	Progress  StoryStatus = "progress"
	Completed StoryStatus = "completed"
)

//Story holds detials of story
type Story struct {
	ID         int
	Title      string      `gorm:"type:varchar(100);unique_index"`
	Status     StoryStatus `sql:"type ENUM('progress', 'completed');default:'progress'"`
	Paragraphs []Paragraph `gorm:"ForeignKey:StoryID"`
}

//Paragraph is linked to a story
//A story can have around configurable paragraph
type Paragraph struct {
	ID        int
	StoryID   int
	Sentences []Sentence `gorm:"ForeignKey:ParagraphID"`
}

//Sentence are linked to paragraph
//A paragraph can have around configurable paragraphs
type Sentence struct {
	ID          uint
	Value       string
	Status      bool
	ParagraphID uint
}

func GetDoubleRelations() {
	story := &Story{}
	db.Preload("Paragraphs").Preload("Paragraphs.Sentences").First(story, 1)

	// It finds the story with the id = 1 and preloads its relationships

	fmt.Printf("%+v\n", story)
}
