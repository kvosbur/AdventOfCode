fn are_safe_values(value_1: i32, value_2: i32, direction: i32) -> bool {
    let difference = value_2 - value_1;
    // handle case where values can only step by 1 - 3
    if difference.abs() < 1 || difference.abs() > 3 {
        return false;
    }

    // handle only increasing/decreasing check
    if difference > 0 && direction == -1 {
        return false;
    }
    if difference < 0 && direction == 1 {
        return false;
    }
    return true;
}

fn line_is_safe(items: &Vec<i32>) -> bool {
    let direction = if items[1] - items[0] > 0 { 1 } else { -1 };
    for value_index in 1..items.len() {
        if !are_safe_values(items[value_index - 1], items[value_index], direction) {
            return false;
        }
    }
    return true;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut safe_count = 0;

    for line in inputs {
        let split: Vec<i32> = line
            .split(" ")
            .map(|item| item.parse::<i32>().unwrap())
            .collect();

        let is_safe = line_is_safe(&split);
        if is_safe {
            safe_count += 1;
            continue;
        }

        //try removing one index at a time
        for delete_index in 0..split.len() {
            let mut cloned = split.clone();
            cloned.remove(delete_index);

            let is_safe = line_is_safe(&cloned);
            if is_safe {
                safe_count += 1;
                break;
            }
        }
    }

    return format!("{}", safe_count);
}
