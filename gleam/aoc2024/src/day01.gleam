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
  let filepath = "./day01_test.txt"
  case simplifile.read(from: filepath) {
    Ok(content) -> parse(content)
    Error(_) -> io.println("file error")
  }
}

fn parse(s: String) -> Nil {
  // io.println("content: \n" <> s)
  let _ =
    string.split(s, on: "\n")
    |> list.map(split_spaces)
    |> list.map(get_edge)
    |> list.map(option.from_result)
    |> option.values
    |> list.map(to_ints)
    |> option.values
    |> io.debug
  io.println("")
}

fn split_spaces(s: String) -> List(String) {
  string.split(s, on: " ")
}

fn get_edge(arr: List(String)) -> Result(#(String, String), AppError) {
  case list.first(arr), list.last(arr) {
    Ok(first), Ok(last) -> Ok(#(first, last))
    _, _ -> Error(ParseError)
  }
}

fn to_ints(t: #(String, String)) -> option.Option(#(Int, Int)) {
  case int.parse(t.0), int.parse(t.1) {
    Ok(a), Ok(b) -> option.Some(#(a, b))
    _, _ -> option.None
  }
}
