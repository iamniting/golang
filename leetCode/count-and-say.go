// https://leetcode.com/problems/count-and-say
// Just sol to the problem, It does not include the I/O part

func countAndSay(n int) string {
    seq := "1"

    for i:=0; i<n-1; i++ {

        nextSeq := ""
        count := 0
        prevChar := seq[0]

        for j:=0; j<len(seq); j++ {

            if prevChar == seq[j] && j == len(seq)-1 {
                nextSeq += string(strconv.Itoa(count+1)) + string(prevChar)
                break
            } else if prevChar != seq[j] && j == len(seq)-1 {
                nextSeq += string(strconv.Itoa(count)) + string(prevChar)
                nextSeq += string(strconv.Itoa(1)) + string(seq[j])
                break
            }

            if prevChar == seq[j] {
                count++
            } else {
                nextSeq += string(strconv.Itoa(count)) + string(prevChar)
                count = 1
            }

            prevChar = seq[j]
        }

        seq = nextSeq
    }
    return seq
}
