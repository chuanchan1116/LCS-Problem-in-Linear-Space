package main

import(
    "bufio"
    "os"
    "strconv"
    "strings"
)

func Max(x []int) (m, pos int) {
    for i, v := range(x) {
        if m < v {
            pos, m = i, v
        }
    }
    return
}

func SubOf(s, substr string) string {
    if strings.Contains(s, substr) {
        return substr
    } else {
        return ""
    }
}

func Common(x, y string) []int {
    c := make([]int, len(y) + 2)
    for i := 0; i < len(x); i++ {
        leftTopScore := 0
        for j := 1; j <= len(y); j++{
            topScore := c[j]
            if x[i] == y[j - 1] {
                c[j] = leftTopScore + 1
            } else {
                c[j], _ = Max([]int{topScore, c[j - 1]})
            }
            leftTopScore = topScore
        }
    }
    return c
}

func CommonReverse(x, y string) []int {
    c := make([]int, len(y) + 2)
    for i := len(x) - 1; i >= 0; i-- {
        rightDownScore := 0
        for j := len(y); j >= 1; j-- {
            downScore := c[j]
            if x[i] == y[j - 1] {
                c[j] = rightDownScore + 1
            } else {
                c[j], _ = Max([]int{downScore, c[j + 1]})
            }
            rightDownScore = downScore
        }
    }
    return c
}

func Lcs(x, y string) string {
    if len(x) == 0 || len(y) == 0 {
        return ""
    }
    if len(x) == 1 {
        return SubOf(y, x)
    }
    if len(y) == 1 {
        return SubOf(x, y)
    }
    xm := len(x) / 2
    c := Common(x[:xm], y)
    cr := CommonReverse(x[xm:], y)
    for i := 0; i < len(cr) - 1; i++ {
        c[i] += cr[i + 1]
    }
    max, k := Max(c)
    if max == 0 {
        return ""
    }
    return Lcs(x[:xm], y[:k]) + Lcs(x[xm:], y[k:])
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    for i := 0; i < n; i++ {
        scanner.Scan()
        x := scanner.Text()
        scanner.Scan()
        y := scanner.Text()
        out.WriteString(Lcs(x, y) + "\n")
    }
}
