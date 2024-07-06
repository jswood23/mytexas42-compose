# Prompt the user for the passphrase
echo "Enter your SSH key passphrase: "
stty -echo
read SSH_PASSPHRASE
stty echo
echo

# Run the Go application, passing the passphrase as an argument
sudo nohup sudo go run . "$SSH_PASSPHRASE" &