// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side
// Just sol to the problem, It does not include the I/O part

func replaceElements(arr []int) []int {

    l := len(arr)

    for i:=0; i<l-1; i++ {

        max := arr[i+1]

        for j:=i+2; j<l; j++ {
            max = int(math.Max(float64(max), float64(arr[j])))
        }
        arr[i] = max
    }

    arr[l - 1] = -1

    return arr
}
