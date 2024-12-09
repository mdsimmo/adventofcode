
let print_files files =
        List.iter (fun x -> if x == -1 then Printf.printf ". " else Printf.printf "%d " x) (files) ;
        print_newline ()

let value_of_char c =
        int_of_char c - int_of_char '0'

let char_of_value n =
        char_of_int (n + int_of_char '0')

let rec append c n list = 
        if n <= 0 then
                list
        else
                append c (n-1) (c :: list)

let repeat x n = 
        append x n []

let parse_in name = 
        let file = open_in name in
        let string = input_line file in
        close_in file ;
        let (seq, _) = String.fold_right (fun c (seq, index) -> 
                (append 
                        (if index mod 2 == 0 then index / 2 else -1) 
                        (value_of_char c) 
                        seq,
                index - 1)
        ) string ([], String.length string - 1) in
        seq

let rec pop_tail list =
        match list with 
        | [] -> ([], None)
        | [x] -> ([], Some x)
        | head :: tail ->
                let (new_tail, x) = (pop_tail tail) in
                (head :: new_tail, x)

let rec set_index list index value =
        match list with
        | [] -> raise (Invalid_argument "Index out of bounds")
        | head :: tail -> 
                if index == 0 then
                        value :: tail
                else
                       head :: set_index tail (index - 1) value
                        

let rec condense files =
        match List.find_index (fun c -> c == -1) files with
        | None -> files
        | Some x -> (
                match pop_tail files with
                | (_, None) -> raise (Invalid_argument "Empty list!")
                | (list, Some c) -> if c == -1 then
                        condense list
                else
                        let new_list = set_index list x c in
                        condense new_list
        )


let find_space files size =
        let rec inner files size remaining index =
                if remaining == 0 then
                        Some (index-size)
                else
                        match files with
                        | [] -> None
                        | x::tail -> if x == -1 then
                                inner tail size (remaining-1) (index+1)
                        else
                                inner tail size size (index+1)
        in
        inner files size size 0

let rec split list n =
        if n < 0 then
                raise (Invalid_argument "split, n < 0")
        else if n == 0 then
                ([], list)
        else
                match list with
                | [] -> raise (Invalid_argument "List too small")
                | x :: tail -> 
                        let (head, tail) = split tail (n-1) in
                        (x :: head, tail)

let rec condense2 files =
        let condense_disk files disk = 
                match List.find_index (fun x -> x == disk) files with
                | None -> raise (Invalid_argument "Disk Not Found")
                | Some index -> (
                        let (pre_disk, post_disk) = split files index in
                        let end_index = List.find_index (fun x -> x != disk) post_disk in
                        let size = match end_index with
                        | None -> List.length post_disk
                        | Some x -> x in

                        let (disk, post_disk) = split post_disk size in
                        match find_space pre_disk size with
                        | None ->
                                files
                        | Some space_index ->
                                let (pre_space, post_space) = split pre_disk space_index in
                                let (_, post_space) = split post_space size in
                                let space = repeat (-1) size in
                                pre_space @ disk @ post_space @ space @ post_disk
                        )
        in
        
        let rec inner_loop files disk =
                if disk < 0 then
                        files
                else
                        inner_loop (condense_disk files disk) (disk-1)
        in
        let max_disk = List.fold_left max 0 files in
        inner_loop files max_disk
        


let checksum files = 
        let rec sum_part index list =
                match list with
                | [] -> 0
                | c :: tail -> index * (if c < 0 then 0 else c) + (sum_part (index+1) tail)
        in sum_part 0 files


let () =
        let files = parse_in "in.txt" in

        let files1 = condense files in
        let sum1 = checksum files1 in
        Printf.printf "Checksum 1: %d" sum1 ;
        print_newline () ;

        let files2 = condense2 files in
        let sum2 = checksum files2 in
        Printf.printf "Checksum 2: %d" sum2 ;
        ()
