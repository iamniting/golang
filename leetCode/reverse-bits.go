// https://leetcode.com/problems/reverse-bits
// Just sol to the problem, It does not include the I/O part

func reverseBits(num uint32) uint32 {
    res := uint32(0)

    for i:=0; i<32; i++ {
        res = (res << 1) | (num & 1)
        num >>= 1
    }
    return res
}