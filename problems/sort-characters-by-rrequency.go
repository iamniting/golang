// https://leetcode.com/problems/sort-characters-by-frequency/
// Just sol to the problem, It does not include the I/O part

func frequencySort(s string) string {

    type key struct { key byte; fre int }

    // create a map with char freq
    m := make(map[byte]int)
    for _, c := range s{ m[byte(c)]++ }

    // copy map into slice and sort it according to values of map
    slice := []key{}
    for k, v := range m {
        slice = append(slice, key{byte(k), v})
    }
    sort.SliceStable(
        slice, func(i, j int) bool { return slice[i].fre > slice[j].fre })

    // create a res string according to freq
    res := ""
    for _,obj := range slice {
        res = res + strings.Repeat(string(obj.key), obj.fre)
    }

    return res
}
