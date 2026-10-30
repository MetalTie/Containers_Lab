package core

import (
    "bufio"
	"fmt"
	"os"
    
    "Containers_Lab/internal/types"
)

const maxEmployees = 200

type Node struct {
    Name string
    Job  types.JobType
    Age  int
    Next int // index into SortSpace; 0 means null
}

type State struct {
    SortByJob [types.JobTypeCount]int     // heads (indices)
    SortSpace [maxEmployees + 1]Node
    Avail     int                   // simple allocator: 1..N
}

func NewState() *State {
    return &State{Avail: 1}
}

// TODO: implement a comparator: (Age asc, Name asc)
func (st *State) less(i, j int) bool {
    // 1. Retrieve the records from your sort space using indices i and j
    recI := st.SortSpace[i]
    recJ := st.SortSpace[j]

    // 2. Compare by Age ascending
    if recI.Age != recJ.Age {
        return recI.Age < recJ.Age
    }

    // 3. If ages are equal, break the tie by Name ascending
    if recI.Name < recJ.Name{
        return true
    }

    return false
}

// TODO: InsertSorted inserts idx into the circular list for its JobType while preserving order.
// Cases: empty ring; new head; middle; tail.
func (st *State) InsertSorted(idx int) { 
    job := st.SortSpace[idx].Job
    head := st.SortByJob[job]

    switch{
    // case #1 for empty ring
    case head == 0:
        st.SortSpace[idx].Next = idx
        st.SortByJob[job] = idx
    // Case #2 for New head
    case st.less(idx, head):
        curr := head
        for st.SortSpace[curr].Next != head{
            curr = st.SortSpace[curr].Next
        }
        st.SortSpace[idx].Next = head
        st.SortSpace[curr].Next = idx
        st.SortByJob[job] = idx
    // Case #3 for Middle and Tail
    default:
        curr := head
        for {
            next := st.SortSpace[curr].Next
            if next == head || (st.less(curr, idx) && st.less(idx, next)){
                st.SortSpace[idx].Next = next
                st.SortSpace[curr].Next = idx
                break
            }
        curr = next
        }
    }

}

// TODO: LoadFromFile_C reads "Name Job Age" triples until EOF, allocates nodes from SortSpace,
// and calls InsertSorted for each record. Reject/skip unknown jobs gracefully.
// LoadFromFile_C reads "Name Job Age" triples until EOF, allocates nodes, and calls InsertSorted.
func (st *State) LoadFromFile_C(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	in := bufio.NewScanner(f)
	in.Split(bufio.ScanWords) // tokens by whitespace

	for in.Scan() {
		name := in.Text()

		if !in.Scan() {
			break
		}
		jobStr := in.Text()

		if !in.Scan() {
			break
		}
		ageStr := in.Text()

		var age int
		if _, err := fmt.Sscanf(ageStr, "%d", &age); err != nil {
			continue // Skip malformed rows
		}

		job, ok := types.ParseJobType(jobStr)
		if !ok {
			continue // Skip unknown jobs gracefully
		}

		if st.Avail > maxEmployees {
			break
		}

		idx := st.Avail
		st.Avail++

		st.SortSpace[idx] = Node{
			Name: name,
			Job:  job,
			Age:  age,
			Next: 0,
		}

		st.InsertSorted(idx)
	}

	return in.Err()
}

// TODO: PrintBuckets prints each non-empty bucket in JobType definition order.
func (st *State) PrintBuckets() { 
    for j := 0; j < types.JobTypeCount; j++ {
        head := st.SortByJob[j]
        if head == 0 {
            continue
        }

        // types.JobType(j).String() is called automatically by %v or %s!
        fmt.Printf("=== Job Bucket: %v ===\n", types.JobType(j))
        
        curr := head
        for {
            node := st.SortSpace[curr]
            fmt.Printf("  -> Name: %s, Age: %d\n", node.Name, node.Age)
            curr = node.Next
            if curr == head {
                break
            }
        }
    }
}