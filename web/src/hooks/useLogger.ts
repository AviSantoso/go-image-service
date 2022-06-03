import winston from "winston";

const logger = winston.createLogger({
  format: winston.format.combine(
    winston.format.colorize(),
    winston.format.json()
  ),
  transports: [],
});

if (import.meta.env.DEV) {
  logger.transports.push(new winston.transports.Console());
}
