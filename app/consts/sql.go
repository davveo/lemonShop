package consts

const (
	SqlOne = "select s.spec_id id,s.spec_name text,  case category_id " +
		"when ? then true else false end selected  from es_specification s left join  es_category_spec " +
		"cs on s.spec_id=cs.spec_id and category_id=? where s.seller_id=0 and s.disabled=1"
	SqlTwo   = "select * from es_specification s inner join es_category_spec cs on s.spec_id = cs.spec_id  where disabled = 1 and category_id = ? and seller_id = ? and spec_name = ? "
	SqlThree = "select s.spec_id,s.spec_name from es_specification s inner join es_category_spec cs on s.spec_id=cs.spec_id where cs.category_id = ? and (s.seller_id=0 or s.seller_id=?)"
	SqlFour  = "select * from es_specification  where disabled = 1 and seller_id = 0 and spec_name = ? and spec_id!=? "
	SqlFive  = "select * from es_category where name = ? and category_id != ? "
)
