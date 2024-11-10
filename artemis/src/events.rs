use crate::{api::message::Message, payload::Payload};

#[derive(Debug)]
pub enum Event {
    Ready,
    MessageCreate(Message),
    GuildCreate,
    TypingStart,
}

impl Event {
    pub fn from_payload(payload: Payload) -> Self {
        let t = payload.t.unwrap();

        match t.as_str() {
            "READY" => Event::Ready,
            "MESSAGE_CREATE" => Event::MessageCreate(serde_json::from_value(payload.d).unwrap()),
            "GUILD_CREATE" => Event::GuildCreate,
            "TYPING_START" => Event::TypingStart,
            _ => panic!("Unknown event: {}", t),
        }
    }
}
