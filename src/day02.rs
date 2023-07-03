use super::inputs;

fn part1(input: &Vec<String>) {
    let mut total_score = 0;
    for pair_str in input {
        if pair_str == "" {
            continue;
        }
        let pair: Vec<i32> = pair_str.split(" ")
            .map(|x| (x.chars().next().unwrap() as i32))
            .collect();
        let left: i32 = pair.get(0).unwrap_or(&0) - 64;
        let right: i32 = pair.get(1).unwrap_or(&0) - 87;
        match right - left {
            2 |-1 => total_score += 0 + right,
            -2| 1 => total_score += 6 + right,
            0 => total_score += 3 + right,
            _ => {}
        }
    }
    println!("| {0: <10} |", "Score");
    println!("| {0: <10} |", "----------");
    println!("| {0: <10} |", total_score);
}


fn part2(input: &Vec<String>) {
    let mut total_score = 0;
    for pair_str in input {
        if pair_str == "" {
            continue;
        }
        let pair: Vec<&str> = pair_str.split(" ").collect();
        let left: i32 = (pair.get(0).unwrap_or(&"A").chars().next().unwrap() as i32) - 64;
        let right: &str = pair.get(1).unwrap_or(&"X");
        match right {
            "X" => total_score += if left - 1 > 0 { left - 1 } else { 3 } ,
            "Y" => total_score += left + 3,
            "Z" => total_score += if left + 1 < 4 { left + 1 } else { 1 },
            _ => {}
        }
    }
    println!("| {0: <10} |", "Score");
    println!("| {0: <10} |", "----------");
    println!("| {0: <10} |", total_score);
}

pub fn solution() {
    let sep: &str = "=====";
    println!("\n{0} DAY 02 {0}", sep);
    let input = inputs::get_input("day02.input");
    println!("[PART 01]:");
    part1(&input);
    println!("[PART 02]:");
    part2(&input);
}
