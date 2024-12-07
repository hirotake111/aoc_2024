import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

const example = "190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"

type Equation {
  Equation(target: Int, nums: List(Int))
}

pub fn start() {
  let assert 3749 = part1(example)
  let assert 11_387 = part2(example)
  let assert Ok(input) = simplifile.read("./data/day07.txt")
  io.println("Part1 -> " <> part1(input) |> int.to_string)
  io.println("Part2 -> " <> part2(input) |> int.to_string)
  Nil
}

fn inner_concat(a: Int, b: Int, n: Int) -> Int {
  case n <= b {
    True -> inner_concat(a, b, n * 10)
    False -> a * n + b
  }
}

fn concat(a: Int, b: Int) -> Int {
  inner_concat(a, b, 10)
}

fn part2(input: String) -> Int {
  let lines = string.split(string.trim(input), "\n")
  let equations = list.map(lines, parse_equation)
  let ops: List(fn(Int, Int) -> Int) = [
    fn(a: Int, b: Int) -> Int { a + b },
    fn(a: Int, b: Int) -> Int { a * b },
    concat,
  ]
  list.fold(equations, 0, fn(sum, eq) { sum + get_calibration_result(eq, ops) })
}

fn part1(input: String) -> Int {
  let lines = string.split(string.trim(input), "\n")
  let equations = list.map(lines, parse_equation)
  let ops: List(fn(Int, Int) -> Int) = [
    fn(a: Int, b: Int) -> Int { a + b },
    fn(a: Int, b: Int) -> Int { a * b },
  ]
  list.fold(equations, 0, fn(sum, eq) { sum + get_calibration_result(eq, ops) })
}

fn get_calibration_result(eq: Equation, ops: List(fn(Int, Int) -> Int)) -> Int {
  let Equation(target, nums) = eq
  let lists =
    list.fold(nums, [], fn(acc, cur) {
      case acc {
        [] -> [cur]
        _ -> {
          list.map(acc, fn(v) { list.map(ops, fn(op) { op(v, cur) }) })
          |> list.flatten
        }
      }
    })
  case list.find(lists, fn(v) { v == target }) {
    Ok(_) -> target
    _ -> 0
  }
}

fn parse_equation(s: String) -> Equation {
  let lines = string.split(s, ": ")
  let assert Ok(left) = list.first(lines)
  let assert Ok(target) = int.parse(left)
  let assert Ok(right) = list.last(lines)
  let nums =
    string.split(right, " ")
    |> list.map(fn(v) {
      let assert Ok(v) = int.parse(v)
      v
    })
  Equation(target: target, nums: nums)
}
