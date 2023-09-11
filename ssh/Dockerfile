# Use Alpine Linux as the base image
FROM alpine:3.18.4

# Install OpenSSH
RUN apk add --no-cache openssh

# Generate host keys
RUN ssh-keygen -A

# Add a user
RUN adduser -D test && echo "test:foobar" | chpasswd

# Expose the SSH port
EXPOSE 22

# Run sshd
CMD ["/usr/sbin/sshd", "-D"]

