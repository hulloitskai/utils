FROM busybox:1.30

# Copy built binary.
COPY ./bread /bin/bread

ENTRYPOINT ["bread"]
