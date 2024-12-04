import gleam/io
import simplifile

pub fn parse() {
  let filepath = "./day01_test.txt"
  case simplifile.read(from: filepath) {
    Ok(v) -> io.println(v)
    Error(_) -> io.println("file error")
  }
}
