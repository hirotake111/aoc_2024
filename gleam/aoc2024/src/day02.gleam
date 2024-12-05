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
  let assert Ok(_raw_data) = simplifile.read("./day02/input.txt")
  let assert 2 = part1(example)
  Nil
}

fn part1(raw_data: String) -> Int {
  let grid =
    raw_data
    |> string.trim
    |> to_grid

  let count =
    grid
    |> list.map(calc)
    |> option.values
    |> list.length
  io.println("Part1 -> " <> count |> int.to_string)
  count
}

fn calc(arr: List(Int)) -> Option(Int) {
  let _increasing =
    arr
    |> io.debug
    |> is_increasing(0, 0)
    |> io.debug
  let _diff =
    arr
    |> list.window_by_2
    |> list.map(fn(pair) { pair.0 - pair.1 })
    |> io.debug
  None
}

fn is_increasing(arr: List(Int), inc: Int, dec: Int) -> Bool {
  case arr {
    [first, second, ..rest] if first < second -> {
      is_increasing(list.append([second], rest), inc + 1, dec)
    }
    [_, second, ..rest] -> {
      is_increasing(list.append([second], rest), inc, dec + 1)
    }
    [] | [_] -> inc > dec
  }
}

fn to_grid(s: String) -> List(List(Int)) {
  string.split(s, "\n")
  |> list.map(to_int_list)
}

fn to_int_list(s: String) -> List(Int) {
  string.split(s, " ")
  |> list.map(int.parse)
  |> list.map(option.from_result)
  |> option.values
}
