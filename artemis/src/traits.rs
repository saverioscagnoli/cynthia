use crate::api::message::Message;

pub trait DiscordHandler {
    async fn ready(&mut self) {}
    async fn message_create(&mut self, _msg: Message) {}
}
