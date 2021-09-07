# Fuck-with-Task-Manager
Using the Windows API to mess with Task Manager in GoLANG

# Info
 When ran you can press a key and it will hide most of the info in Task Manager if its open. Gets the handle of Task Manager via is Window name and then gets the child handle for the main form, then using user32.dll@ShowWindow it hides the form.
 
# Running
 - go build *.go
 - Run as Adminitrator
 - Open Task Manager
 - Press Any Key

# Image
 ![Image](https://i.imgur.com/3W6XQyU.gif)

# Other
Go is a amazing and powerful programming language. If you already haven't, check it out; https://golang.org/

# Donations
<img src="https://blockchain.info/Resources/buttons/donate_64.png"/>
<p align="center">Please Donate To Bitcoin Address: <b>1AEbR1utjaYu3SGtBKZCLJMRR5RS7Bp7eE</b></p>
