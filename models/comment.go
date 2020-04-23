package models

type Comment struct {
	Model
	Title string `json:"title"`
	Content string `json:"content"`
	PostedBy string `json:"posted_by"`
}

func ExistCommentByID(id int)bool {
	var comment Comment
	db.Select("id").Where("id = ?", id).First(&comment)
	return comment.ID > 0
}

func GetCommentCount(maps interface{}) (count int){
	db.Model(&Comment{}).Where(maps).Count(&count)
	return
}

func GetComments(pageNum int, pageSize int, maps interface{}) (comment []Comment){
	db.Where(maps).Offset((pageNum- 1)*pageSize).Limit(pageSize).Find(&comment)
	return
}

func GetComment(id int)(comment Comment){
	db.Where("id = ?", id).First(comment)
	return
}

func AddComment(data map[string]interface{}) bool{
	db.Create(&Comment{
		Title: data["title"].(string),
		Content: data["content"].(string),
		PostedBy: data["posted_by"].(string),
	})
	return true
}


