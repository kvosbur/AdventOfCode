use std::collections::HashMap;

struct StoneCounting {
    step_count: u128,
    stone_map: HashMap<u128, Vec<u128>>,
    stone_count: u128,
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
) -> (u128, &mut StoneCounting) {
    let mut stone_count = 0;
    let mut temp = stone_counting;

    if !temp.stone_map.contains_key(&value) {
        temp.stone_map
            .insert(value, get_stones(value, step_count_step));
    }
    let stones: Vec<u128> = temp.stone_map.get(&value).unwrap().clone();

    if step_count > 59 {
        println!("{} val:{}  stones: {}", step_count, value, stones.len());
    }

    if step_count == step_count_step {
        return (stones.len().try_into().unwrap(), temp);
    }

    for stone in stones {
        let count: u128;
        (count, temp) = get_stone_count(stone, step_count - step_count_step, step_count_step, temp);
        stone_count += count;
    }

    (stone_count, temp)
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let start_values: Vec<u128> = inputs[0]
        .split(" ")
        .map(|val| val.parse().unwrap())
        .collect();

    let mut stone_count = 0;
    let mut stone_counting = StoneCounting {
        step_count: 0,
        stone_map: HashMap::new(),
        stone_count: 0,
    };
    for val in start_values {
        let temp_stone_count: u128;
        (temp_stone_count, _) = get_stone_count(val, 75, 15, &mut stone_counting);
        // let stones = get_stones(val, 75);
        stone_count += temp_stone_count;
    }
    stone_count.to_string()
}
