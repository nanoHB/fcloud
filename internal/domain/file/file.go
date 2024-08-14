package file

type File struct {
	Id             string `bson:"_id,omitempty" json:"id"`
	Name           string `bson:"name" json:"name"`
	IsFolder       bool   `bson:"is_folder" json:"is_folder"`
	Size           int    `bson:"size" json:"size"`
	Path           string `bson:"path" json:"path"`
	ParentFolder   string `bson:"parent_folder" json:"parent_folder"`
	AncestorFolder string `bson:"ancestor_folder" json:"ancestor_folder"`
	CreatedDate    string `bson:"created_date" json:"created_date"`
	UpdatedDate    string `bson:"updated_date" json:"updated_date"`
}
