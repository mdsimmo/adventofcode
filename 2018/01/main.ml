let rec parse_input name =
        let file = open_in name in
        let rec reader () = 
                try 
                        let line = input_line file in
                        let x = int_of_string line in
                        x :: reader ()
                with End_of_file ->  
                        close_in file;
                        []
        in
        reader ()
        
let rec sum_ints ints = 
        match ints with
        | x :: y -> x + sum_ints y
        | [] -> 0

module IntSet = Set.Make(Int)

let rec part_2 ints = 
        let rec first_double ints_rem acc sum = 
                match ints_rem with
                | x :: y -> ( 
                        let new_sum = sum + x in
                        match IntSet.mem new_sum acc with
                        | true -> new_sum
                        | false -> 
                                let new_set = IntSet.add new_sum acc in
                                first_double y new_set new_sum
                )
                | [] -> first_double ints acc sum
        in
        first_double ints IntSet.empty 0

let () = 
        let values = parse_input "in.txt" in
        Printf.printf "Sum: %d\n" (sum_ints values) ;
        Printf.printf "First Duplicate: %d\n" (part_2 values)

