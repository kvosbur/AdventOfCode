fn get_stones(value: u128, step_count: u128) -> Vec<u128> {
    let mut stones = vec![value];
    let mut mutable_step_count = step_count;

    while mutable_step_count > 0 {
        let mut new_stones: Vec<u128> = Vec::new();

        for stone in stones {
            let stone_string = stone.to_string();
            if stone == 0 {
                new_stones.push(1);
            } else if stone_string.len() % 2 == 0 {
                let first_val: u128 = stone_string[..stone_string.len() / 2].parse().unwrap();
                let second_val: u128 = stone_string[stone_string.len() / 2..].parse().unwrap();
                new_stones.push(first_val);
                new_stones.push(second_val);
            } else {
                new_stones.push(stone * 2024);
            }
        }
        stones = new_stones;
        mutable_step_count -= 1;
    }
    stones
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let start_values: Vec<u128> = inputs[0]
        .split(" ")
        .map(|val| val.parse().unwrap())
        .collect();

    let mut stone_count = 0;
    for val in start_values {
        let stones = get_stones(val, 25);
        stone_count += stones.len();
    }
    stone_count.to_string()
}
