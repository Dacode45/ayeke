package algorithms

import(
  "testing"
)

func testEq(a, b []int) bool{
  if len(a) != len(b){
    return false
  }
  for i := range a{
    if a[i] != b[i]{
      return false
    }
  }

  return true
}

func TestInsertionSort(t *testing.T){
  A := []int{9,1,3,2,5,6,8,7,4,0}
  InsertionSort(A)
  B := []int{0,1,2,3,4,5,6,7,8,9}
  if !testEq(B,A) {
    t.Errorf("Insertion Sort Failed: got %v, want %v",A,  B)
  }
}

func TestMaxSubArray(t *testing.T){
  A := []int{13,-3,-25,20,-3,-16,-23,-5,-22,15,-4,7,18,20,-7,12}
  expected_max := 15-4+7+18+20-7+12
  low, high, max := FindMaxSubArray(A)
  if max != expected_max{
    t.Errorf("MaxSubArray Failed: got %v, want %v, values from %v to %v", max, expected_max, low, high)
  }
}

func TestSquareMatrixMultiply(t *testing.T){
  A := Matrix{{1,2},{3,4}}
  B := Matrix{{5,6},{7,8}}
  C := SquareMatrixMultiply(A, B)
  if MatrixEquals(C, Matrix{{19,22},{43,50}}){
    t.Errorf("SquareMatrixMultiply Failed: got %v, want %v", C, Matrix{{19,22},{43,50}})
  }
}
