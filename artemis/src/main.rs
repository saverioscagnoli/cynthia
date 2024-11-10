#[macro_use]
extern crate dotenv_codegen;

use client::Client;
use traits::DiscordHandler;

mod api;
mod client;
mod events;
mod log;
mod opcode;
mod payload;
mod stream;
mod traits;

struct Handler;

impl DiscordHandler for Handler {
    async fn ready(&mut self) {
        println!("Ready!");
    }

    async fn message_create(&mut self, msg: api::message::Message) {
        println!("Message: {:?}", msg);
    }
}

#[tokio::main]
async fn main() {
    dotenv::dotenv().ok();
    let mut client = Client::new().await;

    client.login(dotenv!("TOKEN"), &mut Handler).await;
}
