#[macro_export]
macro_rules! info {
    ($msg:expr) => {
        {
            use chrono::Local;
            use colored::Colorize;
            println!(
                "{}",
                format!(
                    "[{}] [{}] {}",
                    "INFO".blue().bold(),
                    Local::now().format("%Y/%m/%d %H:%M:%S").to_string().green(),
                    $msg
                )
            );
        }
    };
    ($msg:expr, $($arg:tt)*) => {
        {
            use chrono::Local;
            use colored::Colorize;
            println!(
                "{}",
                format!(
                    "[{}] [{}] {}",
                    "INFO".blue().bold(),
                    Local::now().format("%Y/%m/%d %H:%M:%S").to_string().green(),
                    format!($msg, $($arg)*)
                )
            );
        }
    };
}

#[macro_export]
macro_rules! warn {
    ($msg:expr) => {
        {
            use chrono::Local;
            use colored::Colorize;
            eprintln!(
                "{}",
                format!(
                    "[{}] [{}] {}",
                    "WARN".yellow().bold(),
                    Local::now().format("%Y/%m/%d %H:%M:%S").to_string().green(),
                    $msg
                )
            );
        }
    };
    ($msg:expr, $($arg:tt)*) => {
        {
            use chrono::Local;
            use colored::Colorize;
            eprintln!(
                "{}",
                format!(
                    "[{}] [{}] {}",
                    "WARN".yellow().bold(),
                    Local::now().format("%Y/%m/%d %H:%M:%S").to_string().green(),
                    format!($msg, $($arg)*)
                )
            );
        }
    };
}

#[macro_export]
macro_rules! error {
    ($msg:expr) => {
        {
            use chrono::Local;
            use colored::Colorize;
            eprintln!(
                "{}",
                format!(
                    "[{}] [{}] {}",
                    "ERROR".red().bold(),
                    Local::now().format("%Y/%m/%d %H:%M:%S").to_string().green(),
                    $msg
                )
            );
        }
    };
    ($msg:expr, $($arg:tt)*) => {
        {
            use chrono::Local;
            use colored::Colorize;
            eprintln!(
                "{}",
                format!(
                    "[{}] [{}] {}",
                    "ERROR".red().bold(),
                    Local::now().format("%Y/%m/%d %H:%M:%S").to_string().green(),
                    format!($msg, $($arg)*)
                )
            );
        }
    };
}

pub(crate) fn init() {
    std::panic::set_hook(Box::new(|info| {
        if let Some(message) = info.payload().downcast_ref::<&str>() {
            let location = info.location().unwrap();

            error!(
                "[{}] {}",
                format!("{}:{}", location.file(), location.line()).red(),
                message.red()
            );
        } else {
            error!("{}", info);
        }
    }));
}
