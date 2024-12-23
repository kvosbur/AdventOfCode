#[allow(unused_imports)]
use crate::days::day8::part1_solution;

#[test]
fn example() {
    let inputs = "............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"
        .split("\n")
        .map(|temp: &str| -> String { temp.to_string() })
        .collect();
    let result = part1_solution::solve(&inputs);
    assert_eq!(result, "14")
}
