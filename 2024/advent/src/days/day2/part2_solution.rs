use std::collections::HashMap;

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut left: Vec<i32> = Vec::new();
    let mut right: HashMap<i32, i32> = HashMap::new();

    for line in inputs {
        let split: Vec<&str> = line.split(" ").collect();
        left.push(split.first().unwrap().parse().unwrap());
        let right_value: i32 = split.last().unwrap().parse().unwrap();
        right
            .entry(right_value)
            .and_modify(|val| *val += 1)
            .or_insert(1);
    }

    let mut cummulative_score = 0;
    for number in left {
        let right_count = right.entry(number).or_default();
        cummulative_score += number * *right_count;
    }
    return format!("{}", cummulative_score);
}
