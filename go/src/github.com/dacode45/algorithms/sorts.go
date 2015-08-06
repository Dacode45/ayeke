package algorithms

import "fmt"

func InsertionSort(Arr []int){
  fmt.Println("") //so the compiler shuts up

  var i int
  var val int
  for j:= 1; j < len(Arr); j++{
    i = j -1
    val = Arr[j]
    for i >= 0 && Arr[i] > val{
      Arr[i+1] = Arr[i]
      i-= 1
    }
    Arr[i+1] = val
  }
}

func MergeSort(A []int){
  B := make([]int, len(A))
  TopDownSplitMerge(A, B, 0, len(A))
}

func TopDownSplitMerge(A  []int,B []int, begin int, end int){
  if end - begin < 2{
    return
  }

  middle := (end - begin)/2
  TopDownSplitMerge(A, B, begin, middle)
  TopDownSplitMerge(A,B, middle,end)
  TopDownMerge(A,B, begin, middle, end)
  CopyArray(A,B, begin, end)
}

func TopDownMerge(A []int, B []int, begin int, middle int , end int){
  i0 := begin
  i1 := middle

  for j := begin; j < end; j++{
    if(i0 < middle && (i1 >= end || A[i0] <= A[i1])){
      B[j] = A[i0]
      i0 += 1
    }else{
      B[j] = A[i1]
      i1 = i1 + 1;
    }
  }
}

func CopyArray(A []int, B []int, begin int, end int){
  for k := begin; k < end; k++{
    A[k] = B[k]
  }
}

func FindMaxCrossSubArray(A []int, low, high int) (int, int, int){
  mid := (high + low)/2
  var leftSum int
  leftSum = -2147483648 //set to negative infinity
  var sum int
  leftIndex := mid
  for i:= mid; i >= low; i--{

    sum += A[i]
    if sum > leftSum{
      leftSum = sum
      leftIndex = i
    }
  }

  var rightSum int
  rightSum =  -2147483648
  sum = 0
  rightIndex := mid
  for j:= mid+1; j <= high; j++{
    sum += A[j]
    if sum > rightSum{
      rightSum = sum
      rightIndex = j
    }
  }
 //fmt.Printf("Cross Sum %v, %v: %v", low, high, leftSum+rightSum)

  return leftIndex, rightIndex, leftSum+rightSum
}
func _FindMaxSubArray(A []int, low, high int) (int, int, int){
  if low == high {
    return low, high, A[low]
  }
  //fmt.Printf("Testing: %v to %v\n", low, high)
  mid := (high+low)/2

  leftLow, leftHigh, leftSum := _FindMaxSubArray(A, low, mid)
  rightLow, rightHigh, rightSum := _FindMaxSubArray(A, mid+1, high)
  crossLow, crossHigh, crossSum := FindMaxCrossSubArray(A, low, high)
  if leftSum >= rightSum{
    if leftSum >= crossSum {
      //fmt.Printf("Best left is : %v to %v\n", leftLow, leftHigh)
      return leftLow, leftHigh, leftSum
    }
  }else{
    if rightSum >= crossSum{

//fmt.Printf("Best right is : %v to %v\n", rightLow, rightHigh)
      return rightLow, rightHigh, rightSum
    }
  }

  //fmt.Printf("Best cross is : %v to %v\n", crossLow, crossHigh)
  return crossLow, crossHigh, crossSum
}

func FindMaxSubArray(A []int) (int,int,int){
  return _FindMaxSubArray(A, 0, len(A)-1)

}
