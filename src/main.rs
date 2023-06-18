mod day01;
mod day02;

fn main() {
    let day = 2;
    match day {
        1 => day01::solution(),
        2 => {
            day02::solution_part1();
            day02::solution_part2();
        },
        _ => {},
    }
}
