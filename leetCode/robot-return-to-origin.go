// https://leetcode.com/problems/robot-return-to-origin
// Just sol to the problem, It does not include the I/O part

func judgeCircle(moves string) bool {
    UD, LR := 0, 0

    for _, move := range moves {
        if move == 'U' {
            UD++
        } else if move == 'D' {
            UD--
        } else if move == 'L' {
            LR++
        } else if move == 'R' {
            LR--
        }
    }

    return UD == 0 && LR == 0
}
