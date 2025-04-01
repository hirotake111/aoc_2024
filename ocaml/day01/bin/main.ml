let read_lines filename =
  let lines = ref [] in
  (* Get an input channel for given file *)
  let chan = open_in filename in
  try
    while true do
      (* Read a line and add it to the list *)
      let line = input_line chan in
      lines := line :: !lines
    done
  with End_of_file ->
    close_in chan;
    List.rev !lines

let () =
  let filename = "data.txt" in
  let oc = open_out filename in
  output_string oc "line 1\nline 2\nline 3\n";
  close_out oc;

  let line_list = read_lines filename in
  List.iter print_endline line_list
