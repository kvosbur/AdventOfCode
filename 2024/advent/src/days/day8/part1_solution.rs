use std::collections::{HashMap, HashSet};

fn calculate_antinodes(
    towers: HashMap<u8, Vec<(i32, i32)>>,
    max_x: i32,
    max_y: i32,
) -> HashSet<(i32, i32)> {
    let mut antinodes: HashSet<(i32, i32)> = HashSet::new();
    for (_tower_key, radio_towers) in towers {
        if radio_towers.len() < 2 {
            continue;
        }
        for tower_index in 0..radio_towers.len() {
            let main_tower = radio_towers[tower_index];
            for other_tower_index in 0..radio_towers.len() {
                if other_tower_index == tower_index {
                    continue;
                }
                let other_tower = radio_towers[other_tower_index];
                // tower - other tower
                let diff_x = main_tower.0 - other_tower.0;
                let diff_y = main_tower.1 - other_tower.1;
                let new_x = main_tower.0 + diff_x;
                let new_y = main_tower.1 + diff_y;
                if new_x < 0 || new_x > max_x || new_y < 0 || new_y > max_y {
                    continue;
                }
                antinodes.insert((new_x, new_y));
            }
        }
    }

    return antinodes;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut radio_towers: HashMap<u8, Vec<(i32, i32)>> = HashMap::new();
    for x in 0..inputs.len() {
        for y in 0..inputs[0].len() {
            let tower = inputs[x].as_bytes()[y];
            if tower != b'.' {
                radio_towers
                    .entry(tower)
                    .and_modify(|towers| {
                        towers.push((x.try_into().unwrap(), y.try_into().unwrap()));
                    })
                    .or_insert(vec![(x.try_into().unwrap(), y.try_into().unwrap())]);
            }
        }
    }

    let results = calculate_antinodes(
        radio_towers,
        (inputs.len() - 1).try_into().unwrap(),
        (inputs[0].len() - 1).try_into().unwrap(),
    );
    results.len().to_string()
}
