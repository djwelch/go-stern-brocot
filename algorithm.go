package sternbrocot

func search(q float64) (Mn int, Md int) {
    // Initialize two values L and H to 0/1 and 1/0, respectively.
    Ln := 0
    Ld := 1
    Hn := 1
    Hd := 0
    qmin := q - 0.00000001
    qmax := q + 0.00000001
    // Until q is found, repeat the following steps:
    for i := 0; i < 1000; i = i + 1 {
        // * Let L = a/b and H = c/d; compute the mediant M = (a + c)/(b + d).
        Mn = Ln + Hn
        Md = Ld + Hd
        M := float64(Mn) / float64(Md)
        if M < qmin {
            // * If M is less than q, then q is in the open interval (M,H); 
            // replace L by M and continue.
            //fmt.Printf("%d / %d\n", Ln, Ld - Md)
            Ln = Mn
            Ld = Md
        } else if M > qmax {
            // * If M is greater than q, then q is in the open interval (L,M); 
            // replace H by M and continue.
            //fmt.Printf("%d / %d\n", Hn, Md - Hd)
            Hn = Mn
            Hd = Md
        } else {
            // In the remaining case, q = M; terminate the search algorithm.
            break
        }
    }
    return
}



