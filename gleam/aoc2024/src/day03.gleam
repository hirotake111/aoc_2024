import gleam/int
import gleam/io
import gleam/list
import gleam/option.{Some}
import gleam/regexp.{Match}
import simplifile

const example = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

pub fn start() {
  let assert 161 = part1(example)
  let assert Ok(content) = simplifile.read("./data/day03.txt")
  io.println("Part1 -> " <> part1(content) |> int.to_string)
  Nil
}

fn part1(s: String) -> Int {
  let assert Ok(re) = regexp.from_string("mul\\((\\d+),(\\d+)\\)")
  regexp.scan(re, s)
  |> list.fold(0, fn(acc, cur) {
    let assert Match(_, [Some(a), Some(b)]) = cur
    let assert Ok(a) = int.parse(a)
    let assert Ok(b) = int.parse(b)
    acc + a * b
  })
}
