// https://leetcode.com/problems/duplicate-zeros
// Just sol to the problem, It does not include the I/O part

func duplicateZeros(arr []int) {
    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            arr = append(arr[:i+1], arr[i:len(arr)-1]...)
            i++
        }
    }
}
