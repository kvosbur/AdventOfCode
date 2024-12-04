use regex::Regex;

fn get_sum_for_slice(search: &str) -> i32 {
    let re = Regex::new(r"mul\(([0-9]{1,3}),([0-9]{1,3})\)").unwrap();
    let mut sum = 0;
    for m in re.captures_iter(&search) {
        let first_num: i32 = m.get(1).unwrap().as_str().parse().unwrap();
        let second_num: i32 = m.get(2).unwrap().as_str().parse().unwrap();

        sum += first_num * second_num
    }
    sum
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut cumulative_sum = 0;

    for line in inputs {
        println!("----------------line diff --------------");
        let mut index = 0;
        let mut enabled = true;
        while index < line.len() {
            if enabled {
                let next_bad_index = line[index..].find("don't()");
                let next_index = match next_bad_index {
                    None => line.len(),
                    Some(i) => index + i + 7,
                };
                println!(
                    "do sum for {} to {}, {} {}",
                    index,
                    next_index,
                    line.len(),
                    &line[index..next_index]
                );
                cumulative_sum += get_sum_for_slice(&line[index..next_index]);
                index = next_index;
                enabled = false;
            } else {
                let next_good_index = line[index..].find("do()");
                let next_index = match next_good_index {
                    None => line.len(),
                    Some(i) => index + i + 4,
                };
                index = next_index;
                enabled = true;
            }
        }
    }

    cumulative_sum.to_string()
}

// too high 92974346
// original 185797128
