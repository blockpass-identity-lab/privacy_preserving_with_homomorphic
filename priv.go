package main

import (
    "fmt"
    "math/big"
    "math/rand"
    "strconv"
    "os"
)

func get_g_p(size string) (g, p *big.Int) {

    if size == "512" {
        p, _ = new(big.Int).SetString("FCA682CE8E12CABA26EFCCF7110E526DB078B05EDECBCD1EB4A208F3AE1617AE01F35B91A47E6DF63413C5E12ED0899BCD132ACD50D99151BDC43EE737592E17", 16)

        g, _ = new(big.Int).SetString("678471B27A9CF44EE91A49C5147DB1A9AAF244F05A434D6486931D2D14271B9E35030B71FD73DA179069B32E2935630E1C2062354D0DA20A6C416E50BE794CA4", 16)

    }
    if size == "768" {
        p, _ = new(big.Int).SetString("E9E642599D355F37C97FFD3567120B8E25C9CD43E927B3A9670FBEC5D890141922D2C3B3AD2480093799869D1E846AAB49FAB0AD26D2CE6A22219D470BCE7D777D4A21FBE9C270B57F607002F3CEF8393694CF45EE3688C11A8C56AB127A3DAF", 16)

        g, _ = new(big.Int).SetString("30470AD5A005FB14CE2D9DCD87E38BC7D1B1C5FACBAECBE95F190AA7A31D23C4DBBCBE06174544401A5B2C020965D8C2BD2171D3668445771F74BA084D2029D83C1C158547F3A9F1A2715BE23D51AE4D3E5A1F6A7064F316933A346D3F529252", 16)

    }
    if size == "1024" {
        p, _ = new(big.Int).SetString("FD7F53811D75122952DF4A9C2EECE4E7F611B7523CEF4400C31E3F80B6512669455D402251FB593D8D58FABFC5F5BA30F6CB9B556CD7813B801D346FF26660B76B9950A5A49F9FE8047B1022C24FBBA9D7FEB7C61BF83B57E7C6A8A6150F04FB83F6D3C51EC3023554135A169132F675F3AE2B61D72AEFF22203199DD14801C7", 16)

        g, _ = new(big.Int).SetString("F7E1A085D69B3DDECBBCAB5C36B857B97994AFBBFA3AEA82F9574C0B3D0782675159578EBAD4594FE67107108180B449167123E84C281613B7CF09328CC8A6E13C167A8B547C8D28E0A3AE1E2BB3A675916EA37F0BFA213562F1FB627A01243BCCA4F1BEA8519089A883DFE15AE59F06928B665E807B552564014C3BFECF492A", 16)

    }
    return
}

func calcit(i, n int, g, p *big.Int, y []*big.Int) (res *big.Int) {

    res, _ = new(big.Int).SetString("1", 10)

    for j := 0; j <= i-1; j++ {

        gm := y[j]
        res = new(big.Int).Mul(res, gm)
        res = new(big.Int).Mod(res, p)

    }
    for j := i + 1; j < n; j++ {

        gm := y[j]
        gm2 := new(big.Int).ModInverse(gm, p)
        res = new(big.Int).Mul(res, gm2)
        res = new(big.Int).Mod(res, p)

    }
    fmt.Printf(" ")
    return new(big.Int).Mod(res, p)

}
func mult(n int, val []*big.Int, p *big.Int) (r *big.Int) {
    res := val[0]

    for j := 1; j < n; j++ {
        res = new(big.Int).Mul(res, val[j])
    }
    return (new(big.Int).Mod(res, p))

}
func makeG(g *big.Int, val string, p *big.Int) (r *big.Int) {

    v := BigInt(val)

    return (new(big.Int).Exp(g, v, p))

}
func BigInt(val string) (r *big.Int) {
    r, _ = new(big.Int).SetString(val, 10)
    return
}
func getVote(Y *big.Int, x string, p, V *big.Int) (r *big.Int) {
    r = new(big.Int).Mul(new(big.Int).Exp(Y, BigInt(x), p), V)
    return r
}

func main() {


    argCount := len(os.Args[1:])


    n := 5

    x := make([]string, n)
    v := make([]string, n)
    y := make([]*big.Int, n)
    V := make([]*big.Int, n)
    Y := make([]*big.Int, n)
    RegVote := make([]*big.Int, n)


        if (argCount>0) {v[0] = string(os.Args[1])}
        if (argCount>1) {v[1] = string(os.Args[2])}
        if (argCount>2) {v[2] = string(os.Args[3])}
        if (argCount>3) {v[3] = string(os.Args[4])}
        if (argCount>4) {v[4] = string(os.Args[5])}

    g, p := get_g_p("512")

    // Voter random values
    x[0] = strconv.Itoa(rand.Intn(1e12))
    x[1] = strconv.Itoa(rand.Intn(1e12))
    x[2] = strconv.Itoa(rand.Intn(1e12))
    x[3] = strconv.Itoa(rand.Intn(1e12))
    x[4] = strconv.Itoa(rand.Intn(1e12))

    // Voting keys (broadcast by each voter)
    y[0] = makeG(g, x[0], p)
    y[1] = makeG(g, x[1], p)
    y[2] = makeG(g, x[2], p)
    y[3] = makeG(g, x[3], p)
    y[4] = makeG(g, x[4], p)

    // Voter values
    V[0] = makeG(g, v[0], p)
    V[1] = makeG(g, v[1], p)
    V[2] = makeG(g, v[2], p)
    V[3] = makeG(g,v[3], p)
    V[4] = makeG(g, v[4], p)

    // Each voter calculates Y[i]
    Y[0] = calcit(0, 5, g, p, y)
    Y[1] = calcit(1, 5, g, p, y)
    Y[2] = calcit(2, 5, g, p, y)
    Y[3] = calcit(3, 5, g, p, y)
    Y[4] = calcit(4, 5, g, p, y)

    RegVote[0] = getVote(Y[0], x[0], p, V[0])
    RegVote[1] = getVote(Y[1], x[1], p, V[1])
    RegVote[2] = getVote(Y[2], x[2], p, V[2])
    RegVote[3] = getVote(Y[3], x[3], p, V[3])
    RegVote[4] = getVote(Y[4], x[4], p, V[4])

    res0 := mult(5, RegVote, p)

    fmt.Printf("Vote1: %s, Vote2: %s, Vote3: %s, Vote4: %s, Vote5: %s\n", v[0],v[1],v[2],v[3],v[4])

    fmt.Printf("\n\nResult: %s", res0)

    
    for i := 1; i <= 10000; i++ {
        m, _ := new(big.Int).SetString(strconv.Itoa(i), 10)

        gm := new(big.Int).Exp(g, m, p)

        if gm.Cmp(res0) == 0 {
            fmt.Printf("\n\nTotal votes: %d", i)
            break
        }

    }

    fmt.Printf("\n\nVoter 1 (Vote Registration): %s", y[0])
    fmt.Printf("\n\nVoter 2 (Vote Registration): %s", y[1])
    fmt.Printf("\n\nVoter 3 (Vote Registration): %s", y[2])
    fmt.Printf("\n\nVoter 4 (Vote Registration): %s", y[3])
    fmt.Printf("\n\nVoter 5 (Vote Registration): %s", y[4])

    fmt.Printf("\n\nVoter 1 (Vote Value): %s", RegVote[0])
    fmt.Printf("\n\nVoter 2 (Vote Value): %s", RegVote[1])
    fmt.Printf("\n\nVoter 3 (Vote Value): %s", RegVote[2])
    fmt.Printf("\n\nVoter 4 (Vote Value): %s", RegVote[3])
    fmt.Printf("\n\nVoter 5 (Vote Value): %s", RegVote[4])

}

