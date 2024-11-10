use core::panic;
use futures::StreamExt;
use std::{
    sync::{Arc, LazyLock, OnceLock},
    time::Duration,
};
use tokio::sync::mpsc;
use tokio_tungstenite::tungstenite::Message;

use crate::{
    error,
    events::Event,
    info, log,
    opcode::Opcode,
    payload::Payload,
    stream::{Reader, Writer},
    traits::DiscordHandler,
    warn,
};

static TOKEN: OnceLock<Arc<str>> = OnceLock::new();
static HTTP: LazyLock<reqwest::Client> = LazyLock::new(|| reqwest::Client::new());

pub(crate) fn token() -> Arc<str> {
    TOKEN
        .get_or_init(|| panic!("[this is an internal error] Token not set!"))
        .clone()
}

pub(crate) fn http() -> &'static reqwest::Client {
    &HTTP
}

pub struct Client {
    reader: Reader,
    writer_tx: mpsc::Sender<Payload>,
}

impl Client {
    pub async fn new() -> Self {
        log::init();

        warn!("Connecting to the gateway...");

        let (stream, _) = tokio_tungstenite::connect_async("wss://gateway.discord.gg/?v=10")
            .await
            .unwrap();

        let (writer, reader) = stream.split();

        info!("Connected. Using API version {}", 10);

        let (writer_tx, writer_rx) = mpsc::channel(32);

        tokio::spawn(Self::writer_task(Writer::new(writer), writer_rx));

        Self { reader, writer_tx }
    }

    async fn writer_task(mut writer: Writer, mut receiver: mpsc::Receiver<Payload>) {
        while let Some(payload) = receiver.recv().await {
            if let Err(e) = writer.send(payload).await {
                error!("Error while sending payload: {}", e);
            }
        }
    }

    pub async fn login<T: ToString, H: DiscordHandler>(&mut self, token: T, handler: &mut H) {
        // Save the token
        // I fear this is kinda unsafe, will try to find a better way to do this
        // (It will be too late)
        TOKEN.set(Arc::from(token.to_string())).unwrap();

        // Listen for payloads
        while let Some(raw) = self.reader.next().await {
            match raw {
                Ok(Message::Text(json)) => {
                    let opcode = Opcode::parse(&json);

                    match opcode {
                        // This opcode is received after the client identifies itself.
                        // It is used to dispatch any gateway events.
                        // In this case, the `d` field of the payload contains the event data.
                        // Example: the MESSAGE_CREATE event will contain message data.
                        // Thank god for rust enums
                        Opcode::Dispatch(event) => match event {
                            Event::Ready => {
                                info!("Successfully logged in.");
                                handler.ready().await;
                            }

                            Event::MessageCreate(msg) => {
                                handler.message_create(msg).await;
                            }
                            _ => {}
                        },

                        // Discord may send this opcode,
                        // in which case the client must send a heartbeat payload back.
                        Opcode::Heartbeat => {
                            warn!("Received heartbeat request.");
                            self.send_heartbeat().await;
                        }

                        // This is the first opcode that the client receives after connecting.
                        // Here the heartbeat loop needs to be started and the client needs to indentify itself.
                        Opcode::Hello { heartbeat_interval } => {
                            warn!("Identifying...");

                            // Start the heartbeat loop
                            self.start_heartbeat(heartbeat_interval).await;

                            // Identify
                            self.send_identify().await;
                        }

                        // This opcode is received after the client sends a heartbeat.
                        // If not received, the client should reconnect.
                        Opcode::HeartbeatACK => {
                            info!("Heartbeat acknowledged.");
                        }

                        _ => {}
                    }
                }

                Ok(Message::Close(_)) => {
                    warn!("Connection closed by the server");
                    break;
                }

                Ok(_) => {}
                Err(e) => error!("Error while receiving payload: {}", e),
            }
        }
    }

    async fn send_identify(&self) {
        let _ = self.writer_tx.send(Payload::identify()).await;
    }

    async fn send_heartbeat(&self) {
        let _ = self.writer_tx.send(Payload::heartbeat()).await;
    }

    async fn start_heartbeat(&self, interval: u64) {
        let writer_tx = self.writer_tx.clone();

        tokio::spawn(async move {
            loop {
                tokio::time::sleep(Duration::from_millis(interval)).await;
                warn!("Sending heartbeat...");
                let _ = writer_tx.send(Payload::heartbeat()).await;
            }
        });
    }
}
