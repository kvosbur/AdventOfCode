use advent::Config;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let input_filename = "inputs/day1/input.txt";

    let mut should_profile = false;
    if args.len() > 1 && args[1] == "--profile" {
        should_profile = true;
    }
    let config = Config {
        profile: should_profile,
        file_path: input_filename.to_string(),
    };
    dbg!(should_profile, input_filename);

    advent::run(config);
}
