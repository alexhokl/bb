install:
	go install
completion:
	pr completion bash | sudo tee /etc/bash_completion.d/pr
