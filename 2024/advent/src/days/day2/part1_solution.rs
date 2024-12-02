#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut safe_count = 0;

    for line in inputs {
        let split: Vec<i32> = line
            .split(" ")
            .map(|item| item.parse::<i32>().unwrap())
            .collect();
        let direction = if split[1] - split[0] > 0 { 1 } else { -1 };

        let mut is_safe = true;
        for value_index in 1..split.len() {
            let difference = split[value_index] - split[value_index - 1];
            // handle case where values can only step by 1 - 3
            if difference.abs() < 1 || difference.abs() > 3 {
                is_safe = false;
                break;
            }

            // handle only increasing/decreasing check
            if difference > 0 && direction == -1 {
                is_safe = false;
                break;
            }
            if difference < 0 && direction == 1 {
                is_safe = false;
                break;
            }
        }
        if is_safe {
            safe_count += 1;
        }
    }

    return format!("{}", safe_count);
}
