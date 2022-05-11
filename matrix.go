package matrix

import (
	"fmt"
)
type node struct{
	val float64
	right, down *node
}
type Matrix struct{
	rows, cols int
	topleftcorner, startnode, bottomrightcorner *node
}

func matrix(x int, y int) *Matrix{
	m := &Matrix{}
	m.rows = x
	m.cols = y
	for i := 0; i < x; i++{
		for j := 0; j < y; j++{
			temp := &node{}
			if j == 0 {
				if i == 0 {
					m.topleftcorner = temp
					m.startnode = temp
				} else{
					m.startnode.down = temp
					m.startnode = temp
				}
			} else{
				m.bottomrightcorner.right = temp
			}
			m.bottomrightcorner = temp
		}
	}
	return m
}
func (m *Matrix) get(x int, y int) float64{
	if (x <= m.rows) && (y <= m.cols){  
		n := m.topleftcorner
		for i := 1; i < x; i++{
			n = n.down
		}
		for j := 1; j < y; j++{
			n = n.right
		}
		return n.val
	} else{
		fmt.Println("Array bounds exceeded")
		return -1
	}


}
func (m *Matrix) set(x int, y int, value float64){
	if (x <= m.rows) && (y <= m.cols){  
		n := m.topleftcorner
		for i := 1; i < x; i++{
			n = n.down
		}
		for j := 1; j < y; j++{
			n = n.right
		}
		n.val = value
	}
}

func (m *Matrix) sub(r1,c1,r2,c2 int) *Matrix{
	if (r1 <= r2) && (c1 <= c2){
		n := matrix(r2 - r1 + 1, c2 - c1 + 1)
		for i:= 1; i <= n.rows; i++{
			for j := 1; j <= n.cols; j++{
				if (i + r1 - 1 > m.rows) || (j + c1 - 1 > m.cols){
					continue
				}else{
					n.set(i,j,m.get(i+r1-1,j+c1-1))
				}
			}
		}
		return n
	} else{
		panic("Indeces are incorrect")
	}
}

func  size(m *Matrix) *Matrix{
	n := matrix(1,2)
	n.set(1,1,float64(m.rows))
	n.set(1,2,float64(m.cols))
	return n
}
func (m *Matrix) print(){
	n := m.topleftcorner
	for i := 0; i < m.rows; i++{
		temp := n
		for j := 0; j < m.cols; j++{
			fmt.Printf("%+v ",temp.val)
			temp = temp.right
		}
		fmt.Printf("\n")
		n = n.down
	}
	fmt.Println()
}
func  print(m *Matrix){
	n := m.topleftcorner
	for i := 0; i < m.rows; i++{
		temp := n
		for j := 0; j < m.cols; j++{
			fmt.Printf("%+v ",temp.val)
			temp = temp.right
		}
		fmt.Printf("\n")
		n = n.down
	}
	fmt.Println()
}
func (m *Matrix) append(s []float64){
	if m.cols == len(s){
		max := len(s)
		temp := &node{}
		temp.val = s[0]
		m.startnode.down = temp
		m.startnode = temp
		m.bottomrightcorner = temp
		
		     for i := 1; i < max; i++{
			temp1 := &node{}
			temp1.val = s[i]
			m.bottomrightcorner.right = temp1
			m.bottomrightcorner = temp1
		}
	}
	m.rows += 1
}
func trans(m *Matrix) *Matrix{
	n := matrix(m.cols,m.rows)
	for i := 1; i <= n.rows; i++{
		for j := 1; j <= n.cols; j++{
			n.set(i,j,m.get(j,i))
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
func mult(m *Matrix, n *Matrix) *Matrix{
	if (m.cols == n.rows){
		r := matrix(m.rows,n.cols)
		total := 0.0
		for i := 1; i <= r.rows; i++{
			for j := 1; j <= r.cols; j++{
				total = 0
				for k := 1; k <= r.rows; k++{
					total += m.get(i,k) * n.get(k,j)
				}
				r.set(i,j,total)
			}
		} 
		return r
	} else{
		fmt.Println("Matrices are the wrong size")
		r := &Matrix{}
		return r
	}
}
func (m *Matrix) mult(s float64) *Matrix{
	n := matrix(m.rows,m.cols)
	for i := 1; i <= m.rows; i++{
		for j := 1; j <= m.cols; j++{
			n.set(i,j,m.get(i,j) * s)
		}
	}
	return n
}

func (m *Matrix) init(vals ...float64){
	max := len(vals)
	count := 0;
	for i := 1; i <= m.rows; i++{
		for j := 1; (j <= m.cols) && (count < max); j++{
			m.set(i,j,vals[count])
			count++
		}
	}
}
func det(m *Matrix) float64{
	if m.rows == m.cols{

		if m.rows == 2{
			return m.get(1,1)*m.get(2,2) - m.get(1,2)*m.get(2,1)
		}else {
			total := 0.0
			for curr := 1; curr <= m.cols; curr++{
				n := matrix(m.rows - 1, m.rows - 1)
				for i := 2; i <= m.rows; i++{
					col := 1
					for j := 1; j <= m.cols; j++{
						if j == curr{
							continue
						}
						n.set(i-1,col,m.get(i,j))
						col++
					}
				}
				total += m.get(1,curr) * det(n)
			}
			return total
		}
	} else{
		panic("Matrix is not a square")
	}
}
func dot(m,n *Matrix) float64{
	if m.cols == n.cols{
		if m.rows == 1 && n.rows == 1 {
			total := 0.0
			for j := 1; j <= m.cols; j++{
				total += m.get(1,j) * n.get(1,j)
			}
			return total
		}
		panic("Matrices must consist of a single row: dot()")
	}
	panic("Matrices must have the same number of columns: dot()")
}
func cross(m,n *Matrix) *Matrix{
	if m.cols == 3 && n.cols == 3{
		if m.rows == 1 && n.rows == 1 {
			c := matrix(1,3)
			c.set(1,1,(m.get(1,2)*n.get(1,3)) - (m.get(1,3)*n.get(1,2)))
			c.set(1,2,-(m.get(1,1)*n.get(1,3)) + (m.get(1,3)*n.get(1,1)))
			c.set(1,3,(m.get(1,1)*n.get(1,2)) - (m.get(1,2)*n.get(1,1)))
			return c
		}
		panic("Matrices must consist of a single row: cross()")
	}
	panic("Matrices must have three columns: cross()")
}
