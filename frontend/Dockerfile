FROM oven/bun AS builder

WORKDIR /pol-front

COPY . ./

COPY ../.env .

RUN bun install

RUN bun run build

FROM oven/bun

COPY --from=builder /pol-front .

EXPOSE ${POL_FRONT_PORT}

CMD ["bun", "run", "preview"]
