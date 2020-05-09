// https://leetcode.com/problems/destination-city
// Just sol to the problem, It does not include the I/O part

func destCity(paths [][]string) string {

    m := make(map[string]int)

    for _, path := range paths {
        for _, p := range path {
            m[p]++
        }
    }

    for _, path := range paths {
        // start and destination city will be excatly one
        // so check for path[1] == 1
        if m[path[1]] == 1 {
            return path[1]
        }
    }
    return ""
}
