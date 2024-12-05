import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/option
import gleam/result
import gleam/string
import simplifile

pub type AppError {
  ParseError
  FileError(simplifile.FileError)
}

pub fn start() {
  // let filepath = "./day01_test.txt"
  let filepath = "./day01_input.txt"
  io.println("")
  case simplifile.read(from: filepath) {
    Ok(str) -> {
      let str = string.trim(str)
      parse(str)
      similarity_store(str)
    }
    Error(_) -> io.println("file not found: " <> filepath)
  }
}

fn similarity_store(s: String) {
  let lefts = s |> get_lefts
  let rights = s |> get_rights
  let counter =
    list.fold(rights, dict.new(), fn(acc_map, key) {
      let cur = dict.get(acc_map, key) |> result.unwrap(0)
      dict.insert(acc_map, key, cur + 1)
    })
  let total =
    lefts
    |> list.map(fn(v) {
      let n = dict.get(counter, v) |> result.unwrap(0)
      n * v
    })
    |> int.sum
  io.println("Part2 -> " <> int.to_string(total))
}

fn parse(s: String) {
  let lefts = s |> get_lefts |> list.sort(by: int.compare)
  let rights = s |> get_rights |> list.sort(by: int.compare)
  let total =
    list.zip(lefts, rights)
    |> list.map(fn(v) { int.absolute_value(v.1 - v.0) })
    |> int.sum
  io.println("Part1 -> " <> int.to_string(total))
}

fn get_lefts(s: String) -> List(Int) {
  string.split(s, on: "\n")
  |> list.map(split_spaces)
  |> list.map(list.first)
  |> list.map(option.from_result)
  |> option.values
  |> list.map(int.parse)
  |> list.map(option.from_result)
  |> option.values
}

fn get_rights(s: String) -> List(Int) {
  string.split(s, on: "\n")
  |> list.map(split_spaces)
  |> list.map(list.last)
  |> list.map(option.from_result)
  |> option.values
  |> list.map(int.parse)
  |> list.map(option.from_result)
  |> option.values
}

fn split_spaces(s: String) -> List(String) {
  string.split(s, on: " ")
}
