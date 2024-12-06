import gleam/int
import gleam/io
import gleam/list
import gleam/option.{type Option, None}
import gleam/result
import gleam/string
import simplifile

const example = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"

pub fn start() {
  let assert Ok(raw_data) = simplifile.read("./data/day02.txt")
  let assert 2 = part1(example)
  io.println("Part1 -> " <> part1(raw_data) |> int.to_string)
  Nil
}

fn part1(raw_data: String) -> Int {
  raw_data
  |> to_report
  // |> io.debug
  |> list.filter(validate)
  |> list.length
}

fn validate(line: List(Int)) -> Bool {
  let v =
    line
    |> list.window_by_2
    |> list.fold(#(True, True, True), fn(acc, lists) {
      let #(inc, dec, safe_range) = acc
      let #(prev, cur) = lists
      let inc = inc && prev < cur
      let dec = dec && prev > cur
      let safe_range = safe_range && int.absolute_value(prev - cur) <= 3
      #(inc, dec, safe_range)
    })
  { v.0 || v.1 } && v.2
}

fn to_report(s: String) -> List(List(Int)) {
  string.split(string.trim(s), "\n")
  |> list.map(fn(line) {
    string.split(line, " ")
    |> list.map(fn(v) {
      let assert Ok(v) = int.parse(v)
      v
    })
  })
}
