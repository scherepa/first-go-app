package util

import (
	"fmt"
	"math"
	"time"
)

type Post struct {
	Title     string
	Text      string
	CreatedAt time.Time
}

func NewPost(title string, text string) {
	post := Post{title, text, time.Now()}
	fmt.Println("The Post:", post)
	fmt.Printf("%v\n", post)
}

type DateParams struct {
	year  int
	month time.Month
	day   int
}

func intIntoTwoPlacesString(v int) string {
	return fmt.Sprintf("%02d", v)
	// previously used manual for this
	/*if v < 10 {
		return fmt.Sprintf("0%d", v)
	}
	return fmt.Sprintf("%d", v)*/
}

func NewPostClassic(title string, text string) {
	post := Post{
		Title:     title,
		Text:      text,
		CreatedAt: time.Now(),
	}
	// why this approach?
	// it took time to realize format template is a very specific date
	// it is different from other languages I know
	// still keeping it to remind myself the way of my evolution
	fmt.Printf("%v\n", post)
	utc := post.CreatedAt.UTC()
	year := utc.Year()
	month := utc.Month()
	day := utc.Day()
	hour, min, sec := utc.Clock()
	fmt.Printf("Formated date: %d %s %d %s:%s:%s UTC\n", day, month.String(), year, intIntoTwoPlacesString(hour), intIntoTwoPlacesString(min), intIntoTwoPlacesString(sec))
	fmt.Println(
		"The Post Title:",
		post.Title,
		"The Post Text:", post.Text,
		"The Post was created at timestamp:",
		utc,
	)
	// lets try Format again
	// UTC in the example will be just string that has to be added in the end
	// the placeholder should be MST then it will be taken from time.Time object
	fmt.Printf("Formated date: %s\n", utc.Format("02 Jan 2006 15:04:05 UTC"))
	fmt.Printf("Formated date: %s\n", utc.Format("02 01 2006 15:04:05 UTC"))
	fmt.Printf("Formated date most correct: %s\n", utc.Format("02/01/2006 15:04:05 MST"))
	fmt.Printf("Formated date: %s\n", utc.Format("02 01 2006 15:04:05"))
	fmt.Printf("Formated date: %s\n", utc.Format("02-01-2006 15:04:05"))
}

type Shape interface {
	Move(int, int) error
	Area() (float32, error)
}

// challange from course
type Square struct {
	X, Y, Length int
}

func NewSquare(x int, y int, l int) (*Square, error) {
	if l <= 0 {
		return nil, fmt.Errorf("No square can be created, with side length less or equal 0")
	}
	b := Square{
		X:      x,
		Y:      y,
		Length: l,
	}
	return &b, nil
}

func (s *Square) Move(dx int, dy int) error {
	if s == nil {
		return fmt.Errorf("no square provided")
	}
	s.X += dx
	s.Y += dy
	return nil
}

func (s *Square) Area() (float32, error) {
	if s == nil {
		return 0.0, fmt.Errorf("no square provided")
	}
	return 1.0 * float32(s.Length*s.Length), nil
}

// challange from course
type Circle struct {
	X, Y   int
	Radius int
}

func NewCircle(x int, y int, l int) (*Circle, error) {
	if l <= 0 {
		return nil, fmt.Errorf("No circle can be created, with radius length less or equal 0")
	}
	b := Circle{
		X:      x,
		Y:      y,
		Radius: l,
	}
	return &b, nil
}

func (s *Circle) Move(dx int, dy int) error {
	if s == nil {
		return fmt.Errorf("no circle provided")
	}
	s.X += dx
	s.Y += dy
	return nil
}

func (s *Circle) Area() (float32, error) {
	if s == nil {
		return 0.0, fmt.Errorf("no square provided")
	}
	return math.Pi * float32(s.Radius*s.Radius), nil
}
