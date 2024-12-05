import gleam/int
import gleam/io
import gleam/list
import gleam/option
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
    Ok(content) -> parse(content)
    Error(_) -> io.println("file not found: " <> filepath)
  }
}

fn parse(s: String) {
  let s = string.trim(s)
  let firsts =
    string.split(s, on: "\n")
    |> list.map(split_spaces)
    |> list.map(list.first)
    |> list.map(option.from_result)
    |> option.values
    |> list.map(int.parse)
    |> list.map(option.from_result)
    |> option.values
    |> list.sort(by: int.compare)
  let lasts =
    string.split(s, on: "\n")
    |> list.map(split_spaces)
    |> list.map(list.last)
    |> list.map(option.from_result)
    |> option.values
    |> list.map(int.parse)
    |> list.map(option.from_result)
    |> option.values
    |> list.sort(by: int.compare)
  let total =
    list.zip(firsts, lasts)
    |> list.map(fn(v) { int.absolute_value(v.1 - v.0) })
    |> int.sum
  io.println("Part1 -> " <> int.to_string(total))
}

fn split_spaces(s: String) -> List(String) {
  string.split(s, on: " ")
}
