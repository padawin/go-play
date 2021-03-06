package main

import ("fmt"; "time"; "bufio"; "math/rand"; "os"; "strconv"; "strings")
const nbBuckets = 5

func main() {
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    introduction()
    mixBuckets(r1)
    var perlNewPlace int = cheat(r1)

    play(perlNewPlace)
}

func readString() string {
    reader := bufio.NewReader(os.Stdin)
    str, _ := reader.ReadString('\n')
    return str
}

func introduction() {
    fmt.Println("Oye Oye brave people! Cometh and dare face me, the Great Nerlim!")
    fmt.Println("Thou can see in front of me five magical buckets and a magical perl")
    fmt.Println("I shalt put the perl in one of the bucket and,")
    fmt.Println("with the help of my magical hand, I shall swap them and confuse")
    fmt.Println("the most attentive of you!")
    fmt.Println("")
    fmt.Println("Pay Attention! (press return to continue)")
    readString()
}

func mixBuckets(r1 *rand.Rand) {
    var chosenBucket int = r1.Intn(nbBuckets)
    fmt.Println("")
    fmt.Println("I put the perl in the bucket ", chosenBucket)
    fmt.Println("Annnd....")
    fmt.Println("")
    fmt.Println("")
    fmt.Println("I mix!")

    for i := 0; i < 20; i += 1 {
        var b1 int = r1.Intn(nbBuckets)
        var b2 int = r1.Intn(nbBuckets)
        fmt.Println("I swappeth the bucket ", (b1 + 1), " and the bucket ", (b2 + 1))
    }
}

func cheat(r1 *rand.Rand) int {
    return r1.Intn(nbBuckets)
}

func play(perlNewPlace int) {
    fmt.Println("")
    fmt.Println("And now is your turn, art thou the falcon of our age, the mosteth")
    fmt.Println("attentive person in the crowd?")
    fmt.Println("")
    fmt.Println("Art thou able to tell me which bucket (from 1 to ", nbBuckets, ") contains")
    fmt.Println("the mightiest magical perl?")
    fmt.Print("Your choice: ")

    var text string = readString()
    textI, _ := strconv.Atoi(strings.TrimSpace(text))
    for {
        if textI < 1 || textI > nbBuckets {
            fmt.Println("Are you trying to choose a bucket on the stall of my")
            fmt.Println("respected (but still obviously inferior) colleague")
            fmt.Println("the (not so) mightest Fandalg?")
        } else if textI == perlNewPlace {
            fmt.Println("By the seven trolls' beard! You found indeed!")
            break
        } else {
            fmt.Println("Alas as I thoughteth, you did not pay enough attention...")
            fmt.Println("The magical perl was in the bucket ", (perlNewPlace + 1))
            fmt.Println("Please cometh back later and test your luck again!")
            break
        }
    }
}
