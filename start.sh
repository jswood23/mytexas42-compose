# Prompt the user for the passphrase
read -sp "Enter your SSH key passphrase: " SSH_PASSPHRASE

# Run the Go application, passing the passphrase as an argument
sudo nohup sudo go run . "$SSH_PASSPHRASE" &
