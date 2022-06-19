fn main() {
    let mut street_name = String::new();
    println!("Enter the name of the Street you grew up on: ");
    std::io::stdin().read_line(&mut street_name).unwrap();
    let mut pet_name = String::new();
    println!("Enter the name of your pet: ");
    std::io::stdin().read_line(&mut pet_name).unwrap();
    println!("Your band name could be {} {}", street_name.trim(), pet_name.trim())
}
