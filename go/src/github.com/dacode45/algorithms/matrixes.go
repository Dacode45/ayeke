package algorithms

type Matrix [][]int

func MatrixEquals (m, n [][]int) bool{
  if len(m) != len(n){
    return false
  }
  for i, _ := range m{
    if len(m[i]) != len(n[i]){

    return false
    }
    for j, _ := range m{
      if m[i][j] != n[i][j]{
        return false
      }
      }
  }
  return true
}

func SquareMatrixMultiply(A, B Matrix) Matrix{
  n := len(A)
  C := make([][]int, n)

  for i := range n {
    C[i] = make([]int, n)
    for j := range n{
      for k := range n{
        C[i][j] = C[i][j] + A[i][k] + A[k][i]
      }
    }
  }
  return C
}
