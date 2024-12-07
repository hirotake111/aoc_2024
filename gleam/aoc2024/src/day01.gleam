import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub type AppError {
  ParseError
  FileError(simplifile.FileError)
}

const example = "3   4
4   3
2   5
1   3
3   9
3   3
"

pub fn start() {
  let lists = parse_lists(example)
  let assert 11 = part1(lists)
  let assert 31 = part2(lists)
  let assert Ok(raw_data) = simplifile.read("./data/day01.txt")
  io.println("Part1 -> " <> parse_lists(raw_data) |> part1 |> int.to_string)
  io.println("Part2 -> " <> parse_lists(raw_data) |> part2 |> int.to_string)
  Nil
}

fn part2(lists: #(List(Int), List(Int))) -> Int {
  let #(left_list, right_list) = lists
  let freq =
    right_list
    |> list.fold(dict.new(), fn(freq, right) {
      let v = dict.get(freq, right) |> result.unwrap(0)
      dict.insert(freq, right, v + 1)
    })
  left_list
  |> list.fold(0, fn(sum, left) {
    let v = dict.get(freq, left) |> result.unwrap(0)
    sum + { left * v }
  })
}

fn part1(lists: #(List(Int), List(Int))) -> Int {
  let left_list = list.sort(lists.0, int.compare)
  let right_list = list.sort(lists.1, int.compare)
  list.zip(left_list, right_list)
  |> list.fold(0, fn(acc, pair) { acc + int.absolute_value(pair.0 - pair.1) })
}

fn parse_lists(s: String) -> #(List(Int), List(Int)) {
  let s = s |> string.trim
  string.split(s, "\n")
  |> list.fold(#([], []), fn(acc, line) {
    let #(left_list, right_list) = acc
    let arr = string.split(line, "   ")
    let left = list.first(arr) |> result.unwrap("0")
    let right = list.last(arr) |> result.unwrap("0")
    let assert Ok(left) = int.parse(left)
    let assert Ok(right) = int.parse(right)
    #([left, ..left_list], [right, ..right_list])
  })
}
