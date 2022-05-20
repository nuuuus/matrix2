package matrix

import (
	"fmt"
	"math"
)

type Node struct {
	val         float64
	right, down *Node
}
type Mat struct {
	rows, cols                                  int
	topleftcorner, startnode, bottomrightcorner *Node
}

func Matrix(x int, y int) *Mat {
	m := &Mat{}
	m.rows = x
	m.cols = y
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			temp := &Node{}
			if j == 0 {
				if i == 0 {
					m.topleftcorner = temp
					m.startnode = temp
				} else {
					m.startnode.down = temp
					m.startnode = temp
				}
			} else {
				m.bottomrightcorner.right = temp
			}
			m.bottomrightcorner = temp
		}
	}
	return m
}
func (m *Mat) Get(x int, y int) float64 {
	if (x <= m.rows) && (y <= m.cols) {
		n := m.topleftcorner
		for i := 1; i < x; i++ {
			n = n.down
		}
		for j := 1; j < y; j++ {
			n = n.right
		}
		return n.val
	} else {
		panic("Array bounds exceeded")
	}

}
func (m *Mat) Set(x int, y int, value float64) {
	if (x <= m.rows) && (y <= m.cols) {
		n := m.topleftcorner
		for i := 1; i < x; i++ {
			n = n.down
		}
		for j := 1; j < y; j++ {
			n = n.right
		}
		n.val = value
	}
}

func (m *Mat) Sub(r1, c1, r2, c2 int) *Mat {
	if (r1 <= r2) && (c1 <= c2) {
		n := Matrix(r2-r1+1, c2-c1+1)
		for i := 1; i <= n.rows; i++ {
			for j := 1; j <= n.cols; j++ {
				if (i+r1-1 > m.rows) || (j+c1-1 > m.cols) {
					continue
				} else {
					n.Set(i, j, m.Get(i+r1-1, j+c1-1))
				}
			}
		}
		return n
	} else {
		panic("Indeces are incorrect")
	}
}

func Size(m *Mat) *Mat {
	n := Matrix(1, 2)
	n.Set(1, 1, float64(m.rows))
	n.Set(1, 2, float64(m.cols))
	return n
}
func (m *Mat)Rows() int {
	return m.rows
}
func (m *Mat)Cols() int {
	return m.cols
}
func (m *Mat) Print() {
	n := m.topleftcorner
	for i := 0; i < m.rows; i++ {
		temp := n
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%+v ", temp.val)
			temp = temp.right
		}
		fmt.Printf("\n")
		n = n.down
	}
	fmt.Println()
}
func Print(m *Mat) {
	n := m.topleftcorner
	for i := 0; i < m.rows; i++ {
		temp := n
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%+v ", temp.val)
			temp = temp.right
		}
		fmt.Printf("\n")
		n = n.down
	}
	fmt.Println()
}

/*func (m *Mat) Append(s []float64){
	if m.cols == len(s){
		max := len(s)
		temp := &Node{}
		temp.val = s[0]
		m.startnode.down = temp
		m.startnode = temp
		m.bottomrightcorner = temp

		     for i := 1; i < max; i++{
			temp1 := &Node{}
			temp1.val = s[i]
			m.bottomrightcorner.right = temp1
			m.bottomrightcorner = temp1
		}
	}
	m.rows += 1
}*/
func (m *Mat) Append(s *Mat) {
	if (m.cols == s.cols) && (s.rows == 1) {
		max := s.cols
		temp := &Node{}
		temp.val = s.Get(1, 1)
		m.startnode.down = temp
		m.startnode = temp
		m.bottomrightcorner = temp

		for i := 2; i <= max; i++ {
			temp1 := &Node{}
			temp1.val = s.Get(1, i)
			m.bottomrightcorner.right = temp1
			m.bottomrightcorner = temp1
		}
		m.rows += 1
	}
}
func Trans(m *Mat) *Mat {
	n := Matrix(m.cols, m.rows)
	for i := 1; i <= n.rows; i++ {
		for j := 1; j <= n.cols; j++ {
			n.Set(i, j, m.Get(j, i))
		}
	}
	return n
}

/*func (m *matrix) append(s *matrix) {
	if m.cols == s.cols {
		n := m.topleftcorner
		for i := 0; i < n.rows; i++{

		}
	} else{
		fmt.Println("Matrices have different number of columns")
	}
}*/
func Mult(m *Mat, n *Mat) *Mat {
	if m.cols == n.rows {
		r := Matrix(m.rows, n.cols)
		total := 0.0
		for i := 1; i <= r.rows; i++ {
			for j := 1; j <= r.cols; j++ {
				total = 0
				for k := 1; k <= m.cols; k++ {
					total += m.Get(i, k) * n.Get(k, j)
				}
				r.Set(i, j, total)
			}
		}
		return r
	} else {
		fmt.Println("Matrices are the wrong size")
		r := &Mat{}
		return r
	}
}
func Elemult(m, n *Mat) *Mat {
	r := Matrix(m.rows, m.cols)
	if (m.rows == n.rows) && (m.cols == n.cols) {
		sm := m.topleftcorner
		sn := n.topleftcorner
		sr := r.topleftcorner
		tm := m.topleftcorner
		tn := n.topleftcorner
		tr := r.topleftcorner
		for i := 0; i < m.rows; i++ {
			for j := 0; j < m.cols; j++ {
				tr.val = tm.val * tn.val
				tr = tr.right
				tm = tm.right
				tn = tn.right
			}
			sm = sm.down
			sn = sn.down
			sr = sr.down
			tm = sm
			tn = sn
			tr = sr
		}
	} else {
		fmt.Println("ERROR: Matrices are different sizes: Elemult")
	}
	return r
}
func (m *Mat) Mult(s float64) *Mat {
	n := Matrix(m.rows, m.cols)
	for i := 1; i <= m.rows; i++ {
		for j := 1; j <= m.cols; j++ {
			n.Set(i, j, m.Get(i, j)*s)
		}
	}
	return n
}

func (m *Mat) Init(vals ...float64) {
	max := len(vals)
	count := 0
	for i := 1; i <= m.rows; i++ {
		for j := 1; (j <= m.cols) && (count < max); j++ {
			m.Set(i, j, vals[count])
			count++
		}
	}
}
func Det(m *Mat) float64 {
	if m.rows == m.cols {

		if m.rows == 2 {
			return m.Get(1, 1)*m.Get(2, 2) - m.Get(1, 2)*m.Get(2, 1)
		} else {
			total := 0.0
			for curr := 1; curr <= m.cols; curr++ {
				n := Matrix(m.rows-1, m.rows-1)
				for i := 2; i <= m.rows; i++ {
					col := 1
					for j := 1; j <= m.cols; j++ {
						if j == curr {
							continue
						}
						n.Set(i-1, col, m.Get(i, j))
						col++
					}
				}
				total += m.Get(1, curr) * Det(n)
			}
			return total
		}
	} else {
		panic("Matrix is not a square")
	}
}
func Norm(m *Mat) float64 {
	return math.Sqrt(Dot(m, m))
}
func Dot(m, n *Mat) float64 {
	if m.cols == n.cols {
		if m.rows == 1 && n.rows == 1 {
			total := 0.0
			for j := 1; j <= m.cols; j++ {
				total += m.Get(1, j) * n.Get(1, j)
			}
			return total
		}
		panic("Matrices must consist of a single row: dot()")
	}
	panic("Matrices must have the same number of columns: dot()")
}
func Cross(m, n *Mat) *Mat {
	if m.cols == 3 && n.cols == 3 {
		if m.rows == 1 && n.rows == 1 {
			c := Matrix(1, 3)
			c.Set(1, 1, (m.Get(1, 2)*n.Get(1, 3))-(m.Get(1, 3)*n.Get(1, 2)))
			c.Set(1, 2, -(m.Get(1, 1)*n.Get(1, 3))+(m.Get(1, 3)*n.Get(1, 1)))
			c.Set(1, 3, (m.Get(1, 1)*n.Get(1, 2))-(m.Get(1, 2)*n.Get(1, 1)))
			return c
		}
		panic("Matrices must consist of a single row: cross()")
	}
	panic("Matrices must have three columns: cross()")
}
func (m *Mat) Copy() *Mat {
	n := Matrix(m.rows, m.cols)
	for i := 1; i <= m.rows; i++ {
		for j := 1; j <= m.cols; j++ {
			n.Set(i, j, m.Get(i, j))
		}
	}
	return n
}

func Add(m ...*Mat) *Mat {
	cont := 1
	max := len(m)
	rows := m[0].rows
	cols := m[0].cols
	for c := 0; c < max; c++ {
		if (rows != m[c].rows) || (cols != m[c].cols) {
			cont = 0
			c = max
		}
	}
	if cont == 1 {
		r := Matrix(rows, cols)
		for c := 0; c < max; c++ {
			sr := r.topleftcorner
			tr := r.topleftcorner
			s := m[c].topleftcorner
			t := m[c].topleftcorner
			for i := 0; i < rows; i++ {
				t = s
				tr = sr
				for j := 0; j < cols; j++ {
					tr.val += t.val
					t = t.right
					tr = tr.right
				}
				s = s.down
				sr = sr.down
			}
		}
		return r
	} else {
		panic("Matrices aren't the same size: Add()")
	}
}
func (m *Mat) Add(s float64) *Mat {
	r := Matrix(m.rows, m.cols)
	sm := m.topleftcorner
	tm := m.topleftcorner
	sr := r.topleftcorner
	tr := r.topleftcorner
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			tr.val = tm.val + s
			tr = tr.right
			tm = tm.right
		}
		sr = sr.down
		sm = sm.down
		tr = sr
		tm = sm
	}
	return r
}
func Span(s, interval, f float64) *Mat {
	m := &Mat{}
	m.topleftcorner = &Node{}
	m.startnode = m.topleftcorner
	m.bottomrightcorner = m.startnode
	m.topleftcorner.val = s
	m.rows = 1
	m.cols = 1
	val := s
	for val < f {
		tail := &Node{}
		val += interval
		tail.val = val
		m.bottomrightcorner.right = tail
		m.bottomrightcorner = tail
		m.cols += 1
	}
	if m.bottomrightcorner.val > f {
		m.cols -= 1
	}
	return m
}
