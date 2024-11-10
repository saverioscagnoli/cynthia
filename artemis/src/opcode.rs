use crate::{events::Event, payload::Payload};

#[derive(Debug)]
pub enum Opcode {
    Dispatch(Event),
    Heartbeat,
    Identify,
    PresenceUpdate,
    VoiceStateUpdate,
    Resume,
    Reconnect,
    RequestGuildMembers,
    InvalidSession,
    Hello { heartbeat_interval: u64 },
    HeartbeatACK,
    RequestSoundboardSounds,
}

impl From<Opcode> for u8 {
    fn from(opcode: Opcode) -> u8 {
        match opcode {
            Opcode::Dispatch(_) => 0,
            Opcode::Heartbeat => 1,
            Opcode::Identify => 2,
            Opcode::PresenceUpdate => 3,
            Opcode::VoiceStateUpdate => 4,
            Opcode::Resume => 6,
            Opcode::Reconnect => 7,
            Opcode::RequestGuildMembers => 8,
            Opcode::InvalidSession => 9,
            Opcode::Hello { .. } => 10,
            Opcode::HeartbeatACK => 11,
            Opcode::RequestSoundboardSounds => 31,
        }
    }
}

impl Opcode {
    pub fn parse(json: &str) -> Self {
        let payload: Payload = serde_json::from_str(json).unwrap();

        match payload.op {
            0 => Opcode::Dispatch(Event::from_payload(payload)),
            1 => Opcode::Heartbeat,
            2 => Opcode::Identify,
            3 => Opcode::PresenceUpdate,
            4 => Opcode::VoiceStateUpdate,
            6 => Opcode::Resume,
            7 => Opcode::Reconnect,
            8 => Opcode::RequestGuildMembers,
            9 => Opcode::InvalidSession,
            10 => Opcode::Hello {
                heartbeat_interval: payload.d["heartbeat_interval"].as_u64().unwrap(),
            },
            11 => Opcode::HeartbeatACK,
            31 => Opcode::RequestSoundboardSounds,
            _ => panic!("Unknown opcode: {}", payload.op),
        }
    }
}
