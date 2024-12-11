use std::collections::HashMap;

struct StoneCounting {
    stone_map: HashMap<u128, Vec<u128>>,
    stone_count_map: HashMap<u128, u128>,
}

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

fn get_stone_count(
    value: u128,
    step_count: u128,
    step_count_step: u128,
    stone_counting: &mut StoneCounting,
) -> u128 {
    let mut stone_count = 0;
    let stone_count_key: u128 = value * 100 + step_count;

    if stone_counting
        .stone_count_map
        .contains_key(&stone_count_key)
    {
        return stone_counting
            .stone_count_map
            .get(&stone_count_key)
            .unwrap()
            .clone();
    }

    if !stone_counting.stone_map.contains_key(&value) {
        stone_counting
            .stone_map
            .insert(value, get_stones(value, step_count_step));
    }

    let stones: Vec<u128> = stone_counting.stone_map.get(&value).unwrap().clone();

    if step_count == step_count_step {
        return stones.len().try_into().unwrap();
    }

    for stone in stones {
        stone_count += get_stone_count(
            stone,
            step_count - step_count_step,
            step_count_step,
            stone_counting,
        );
    }

    stone_counting
        .stone_count_map
        .insert(stone_count_key, stone_count);

    stone_count
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let start_values: Vec<u128> = inputs[0]
        .split(" ")
        .map(|val| val.parse().unwrap())
        .collect();

    let mut stone_count = 0;
    let mut stone_counting = StoneCounting {
        stone_map: HashMap::new(),
        stone_count_map: HashMap::new(),
    };
    for val in start_values {
        stone_count += get_stone_count(val, 75, 5, &mut stone_counting);
    }
    stone_count.to_string()
}
