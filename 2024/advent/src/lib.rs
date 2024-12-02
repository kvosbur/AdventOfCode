mod days;

use days::day2::part2_solution;
use std::fs;

pub struct Config {
    pub profile: bool,
    pub file_path: String,
}

impl Config {
    pub fn read_file_to_vector_lines(&self) -> Vec<String> {
        let mut result = Vec::new();

        for line in fs::read_to_string(&self.file_path)
            .expect("Should have file to read.")
            .lines()
        {
            result.push(line.to_string())
        }

        result
    }
}

pub fn run(config: Config) {
    let lines = config.read_file_to_vector_lines();
    let sol = part2_solution::solve(&lines);
    println!("Solution is: {}", sol);
}
