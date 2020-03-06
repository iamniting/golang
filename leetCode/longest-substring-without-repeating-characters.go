// https://leetcode.com/problems/longest-substring-without-repeating-characters
// Just sol to the problem, It does not include the I/O part

func lengthOfLongestSubstring(s string) int {
    i := 0
    j := 0
    max := 0

    /*
    keep left and right two pointers like a sliding window.
    set is not going to keep duplicate characters.
    when ever adding any char in the set get the max and 
    change the value of max accordingly.
    */

    type void struct {}
    set := make(map[byte]void)

    for j < len(s) {
        // if char is already present in set
        if _, ok := set[s[j]]; ok {
            // delete element from the set and increment the left ptr
            delete(set, s[i])
            i++
        // if char is not present in set
        } else {
            // add element in the set and increment the right ptr
            set[s[j]] = void{}
            j++
            max = int(math.Max(float64(len(set)), float64(max)))
        }
    }
    return max
}
