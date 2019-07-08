package mysql

type TestTable struct {
	Id   int64  `orm:"auto;pk" json:"-"` // not out put json
	Name string `json:"query" orm:"size(32);description(名称)"`
}

func (this *TestTable) TableName() string {
	return "test_table"
}

func (this *TestTable) Read(cols ...string) (err error) {
	if err = m_orm.Read(this, cols...); err == nil {
		return nil
	}
	return err
}

func (this *TestTable) Insert() (int64, error) {
	return m_orm.Insert(this)
}

func (this *TestTable) Update(cols ...string) (int64, error) {
	return m_orm.Update(this, cols...)
}

func (this *TestTable) Delete() (int64, error) {
	return m_orm.Delete(this)
}
