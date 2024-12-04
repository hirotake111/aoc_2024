import argv
import day01
import gleam/io

pub fn main() {
  case argv.load().arguments {
    ["day01"] -> day01.parse()
    _ -> io.println("Usage: aoc <day>")
  }
}
