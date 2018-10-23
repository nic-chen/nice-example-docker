package dao

func NewMemberDao() *Tbl {
	cols := make(map[string]string);
	m := &Tbl{
		Name: "member",
		Key: "id",
		Cols: cols,
	}

	return m;
}