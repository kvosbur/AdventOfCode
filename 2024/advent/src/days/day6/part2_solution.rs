use std::collections::HashMap;

fn valid_configuration(
    split_pages: &Vec<i32>,
    mappings: &mut HashMap<i32, Vec<i32>>,
) -> Option<usize> {
    let mut seen_pages: Vec<i32> = vec![];
    for index in 0..split_pages.len() {
        let page = split_pages[index];
        let bad_pages: &Vec<i32> = mappings.entry(page).or_default();

        let has_seen_bad_page = bad_pages.iter().any(|item| seen_pages.contains(item));
        if has_seen_bad_page {
            return Some(index);
        }
        seen_pages.push(page);
    }
    return None;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut mappings: HashMap<i32, Vec<i32>> = HashMap::new();

    let split_index = inputs.iter().position(|r| r == "").unwrap();
    for setup in &inputs[..split_index] {
        let split_line: Vec<&str> = setup.split("|").collect();
        let first_val: i32 = split_line[0].parse().unwrap();
        let second_val: i32 = split_line[1].parse().unwrap();

        mappings
            .entry(first_val)
            .and_modify(|v| v.push(second_val))
            .or_insert(vec![second_val]);
    }

    let mut cumullative_sum = 0;
    for manual_updates in &inputs[split_index + 1..] {
        let mut split_pages: Vec<i32> = manual_updates
            .split(",")
            .map(|page| page.parse::<i32>().unwrap())
            .collect();

        let initial_bad = valid_configuration(&split_pages, &mut mappings).is_some();
        if initial_bad {
            let mut bad_update = true;
            while bad_update {
                let bad_index = valid_configuration(&split_pages, &mut mappings);
                bad_update = bad_index.is_some();
                if bad_update {
                    let bad_index_val = bad_index.unwrap();
                    let temp = split_pages[bad_index_val];
                    split_pages[bad_index_val] = split_pages[bad_index_val - 1];
                    split_pages[bad_index_val - 1] = temp;
                }
            }
            let middle_index = split_pages.len() / 2;
            cumullative_sum += split_pages[middle_index];
        }
    }

    cumullative_sum.to_string()
}
