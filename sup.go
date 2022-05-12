package matrix

func (m *Mat) Add(n *Mat){
	if (m.rows == n.rows) && (m.cols == n.cols){
		r := Matrix(m.rows,m.cols)
		for i := 1; i <= r.rows; i++{
			for j:=1; j <= r.cols; j++{
				r.Set(i,j,(m.Get(i,j)+n.Get(i,j)))
			}
		}
		return r
	}else{
		panic("Matrices must be the same size: Add()")
	}
}
