package dao

func NewMemberDao() *Tbl {
	m := &Tbl{
		Name: "member",
		Key: "id",
		//Cols: new(map[string]interface{}),
	}

	return m;
}