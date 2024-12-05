import argv
import day01
import day02
import gleam/io

pub fn main() {
  case argv.load().arguments {
    ["day01"] -> day01.start()
    ["day02"] -> day02.start()
    _ -> io.println("Usage: aoc <day>")
  }
}
