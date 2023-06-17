mod day01;
mod day02;

fn main() {
    let day = 1;
    match day {
        1 => day01::solution(),
        2 => day02::solution(),
        _ => {},
    }
}
