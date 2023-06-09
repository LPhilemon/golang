type Track struct {
	Title string
	Artist string
	Album string
	Year
	int
	Length time.Duration
	}
	var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	func length(s string) time.Duration {
d, err := time.ParseDuration(s)
if err != nil {
panic(s)
}
return d
}

func printTracks(tracks []*Track) {
const format = "%v\t%v\t%v\t%v\t%v\t\n"
tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
for _, t := range tracks {
fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
}
tw.Flush() // calculate column widths and print table
}

sort.Sort(customSort{tracks, func(x, y *Track) bool {
if x.Title != y.Title {
return x.Title < y.Title
}
if x.Year != y.Year {
return x.Year < y.Year
}
if x.Length != y.Length {
return x.Length < y.Length
}
return false
}})

values := []int{3, 1, 4, 1}
fmt.Println(sort.IntsAreSorted(values)) // "false"
sort.Ints(values)
fmt.Println(values)
// "[1 1 3 4]"
fmt.Println(sort.IntsAreSorted(values)) // "true"
sort.Sort(sort.Reverse(sort.IntSlice(values)))
fmt.Println(values)
// "[4 3 1 1]"
fmt.Println(sort.IntsAreSorted(values)) // "false"