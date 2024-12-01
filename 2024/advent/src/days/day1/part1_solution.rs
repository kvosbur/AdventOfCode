pub fn solve(inputs: &Vec<String>) -> String {
    let mut left: Vec<i32> = Vec::new();
    let mut right: Vec<i32> = Vec::new();

    for line in inputs {
        let split: Vec<&str> = line.split(" ").collect();
        left.push(split.first().unwrap().parse().unwrap());
        right.push(split.last().unwrap().parse().unwrap());
    }

    left.sort();
    right.sort();
    let mut cummulative_distance = 0;
    for i in 0..left.len() {
        let distance = left[i] - right[i];
        cummulative_distance += distance.abs();
    }
    return format!("{}", cummulative_distance);
}
