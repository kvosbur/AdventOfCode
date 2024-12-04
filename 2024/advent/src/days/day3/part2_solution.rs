use regex::Regex;

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut cumulative_sum = 0;
    let mut enabled = true;

    for line in inputs {
        let re = Regex::new(r"mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)").unwrap();

        for m in re.captures_iter(&line) {
            let statement = m.get(0).unwrap().as_str();
            if statement == "do()" {
                enabled = true
            } else if statement == "don't()" {
                enabled = false
            } else if enabled {
                println!("{}", statement);
                let first_num: i32 = m.get(1).unwrap().as_str().parse().unwrap();
                let second_num: i32 = m.get(2).unwrap().as_str().parse().unwrap();

                cumulative_sum += first_num * second_num
            }
        }
    }

    cumulative_sum.to_string()
}
