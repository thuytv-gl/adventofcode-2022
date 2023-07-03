use super::inputs;

fn part1(input: &Vec<String>) {
    let mut inventory: Vec<u32> = vec![];
    let mut counter: u32 = 0;
    for s in input {
        let num = s.trim().parse::<u32>();
        match num {
            Ok(n) => counter += n,
            Err(e) => match e.kind() {
                std::num::IntErrorKind::Empty => {
                    inventory.push(counter);
                    counter = 0;
                },
                _ => {},
            }
        }
    }

    inventory.sort_by(|a, b| b.cmp(a));

    let top1: u32 = *inventory.get(0).unwrap_or(&0);
    println!("| {0: <10} | {1: <10} |", "Name", "Value");
    println!("| {0: <10} | {1: <10} |", "----------",  "----------");
    println!("| {0: <10} | {1: <10} |", "Top 1", top1);
}

fn part2(input: &Vec<String>) {
    let mut inventory: Vec<u32> = vec![];
    let mut counter: u32 = 0;
    for s in input {
        let num = s.trim().parse::<u32>();
        match num {
            Ok(n) => counter += n,
            Err(e) => match e.kind() {
                std::num::IntErrorKind::Empty => {
                    inventory.push(counter);
                    counter = 0;
                },
                _ => {},
            }
        }
    }

    inventory.sort_by(|a, b| b.cmp(a));

    let top1: u32 = *inventory.get(0).unwrap_or(&0);
    let top3: u32 = inventory.get(0..3).unwrap_or(&[]).iter().sum();
    println!("| {0: <10} | {1: <10} |", "Name", "Value");
    println!("| {0: <10} | {1: <10} |", "----------",  "----------");
    println!("| {0: <10} | {1: <10} |", "Top 1", top1);
    println!("| {0: <10} | {1: <10} |", "Top 3", top3);
}

pub fn solution() {
    let input = inputs::get_input("day01.input");
    let sep: &str = "=====";
    println!("\n{0} DAY 01 {0}", sep);
    println!("[PART 01]:");
    part1(&input);
    println!("[PART 02]:");
    part2(&input);
}

