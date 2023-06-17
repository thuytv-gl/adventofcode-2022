use std::fs;

pub fn solution() {
    let chars: String = fs::read_to_string("src/day01.input").unwrap().parse().unwrap();
    let input: Vec<&str> = chars.split("\n").collect();
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

    println!(
        "| {0: <10} | {1: <10} |",
        "Name", "Value",
    );
    println!("| {0: <10} | {1: <10} |", "----------",  "----------");
    if let Some(top1) = inventory.get(0) {
        println!("| {0: <10} | {1: <10} |", "Top 1", top1);
    }

    let mut top3 = 0;
    for i in 0..3 {
        if let Some(calories) = inventory.get(i) {
            top3 += calories;
        }
    }
    println!("| {0: <10} | {1: <10} |", "Top 3", top3);
}


