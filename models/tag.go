package models

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
/*
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("AddDt", time.Now().Unix())
	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdateDt", time.Now().Unix())
	return nil
}
*/
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagCount(maps interface{}) (count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false

}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:name,
		State:state,
		CreatedBy:createdBy,
	})
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true

}

func EditTag(id int, data interface{}) bool  {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true

}

func ClearAllTag() bool {
	db.Unscoped().Where("deleted_dt != ?", 0).Delete(&Tag{})
	return true
	
}