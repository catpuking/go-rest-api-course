package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our commnet service
type Service struct {
	DB *gorm.DB
}

// Comment - Struct of our comment content
type Comment struct {
	gorm.Model
	Slug	string
	Body	string
	Author  string
}

//CommentService - the interface for our comment service
type CommentService interface{
	GetComment(ID uint)(Comment, error)
	GetCommentsBySlug(slug string)([]Comment, error)
	PostComment(comment Comment)(Comment, error)
	UpdateComment(ID uint, newComment Comment)(Comment, error)
	DeleteComment(ID uint) error
	GetAllComments([]Comment, error)
}

//NewService - Returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - takes in an ID and returns a comment
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//GetCommentsBySlug - takes in a slug and returns an slice of comments
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug) ; result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

//PostComment - adds a new comment to the DB
func (s *Service) PostComment (comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//UpdateComment - takes in an ID & new comment and returns new comment
func (s *Service) UpdateComment (ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// DeleteComment - Takes ID and deletes from DB
func (s *Service) DeleteComment(ID uint) error {
	if results := s.DB.Delete(&Comment{}, ID); results.Error != nil {
		return results.Error
	}
	return nil
}

func (s *Service) GetAllComments () ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil{
		return []Comment{}, result.Error
	}
	return comments, nil
}