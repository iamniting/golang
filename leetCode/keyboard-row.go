// https://leetcode.com/problems/keyboard-row
// Just sol to the problem, It does not include the I/O part

func findWords(words []string) []string {
    m := map[rune]int {
        'q':1,'w':1,'e':1,'r':1,'t':1,'y':1,'u':1,'i':1,'o':1,'p':1,
        'a':2,'s':2,'d':2,'f':2,'g':2,'h':2,'j':2,'k':2,'l':2,
        'z':3,'x':3,'c':3,'v':3,'b':3,'n':3,'m':3,
    }
    res := []string{}

    for _, word := range words {

        row := m[rune(word[0] | ' ')]
        isSameRow := true

        for _, c := range word {
            // if row is not as a first char
            if m[c | ' '] != row {
                isSameRow = false
                break
            }
        }
        // if rows was same append the word in res
        if isSameRow {
            res = append(res, word)
        }
    }
    return res
}
