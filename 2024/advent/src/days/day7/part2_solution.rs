fn combinations(
    target_value: u64,
    current_value: u64,
    current_index: usize,
    values: &Vec<u64>,
) -> u32 {
    if current_index == values.len() - 1 {
        return if current_value == target_value { 1 } else { 0 };
    }

    let mut combo_counts = 0;
    let next_index = current_index + 1;
    combo_counts += combinations(
        target_value,
        current_value + values[next_index],
        next_index,
        values,
    );
    combo_counts += combinations(
        target_value,
        current_value * values[next_index],
        next_index,
        values,
    );
    let mut concat_value = current_value.to_string();
    concat_value.push_str(&values[next_index].to_string());
    combo_counts += combinations(
        target_value,
        concat_value.parse().unwrap(),
        next_index,
        values,
    );
    return combo_counts;
}

#[allow(dead_code)]
pub fn solve(inputs: &Vec<String>) -> String {
    let mut calibration_result = 0;
    for line in inputs {
        let split_index = line.find(":").unwrap();
        let target_value: u64 = line[..split_index].parse().unwrap();
        let other_values: Vec<u64> = line[split_index + 2..]
            .split(" ")
            .map(|val| val.parse::<u64>().unwrap())
            .collect();

        let combo_counts = combinations(target_value, other_values[0], 0, &other_values);
        // println!("target: {} combo_counts:{}", target_value, combo_counts);
        if combo_counts > 0 {
            calibration_result += target_value;
        }
    }
    calibration_result.to_string()
}
