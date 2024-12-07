import argv
import day01
import day02
import day03
import gleam/io

pub fn main() {
  case argv.load().arguments {
    ["day01"] -> day01.start()
    ["day02"] -> day02.start()
    ["day03"] -> day03.start()
    _ -> io.println("Usage: aoc <day>")
  }
}
